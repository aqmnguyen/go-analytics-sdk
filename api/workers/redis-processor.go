package workers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"analytics-api/types"

	"github.com/redis/go-redis/v9"
)

type RedisProcessor struct {
	redisClient *redis.Client
	db          *sql.DB
	ctx         context.Context
}

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379", // Redis server address
		Password:     "",               // No password for local development
		DB:           0,                // Default DB
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Warning: Could not connect to Redis: %v", err)
		log.Println("Redis connection failed - events will not be queued")
		return nil
	}

	log.Println("Successfully connected to Redis")
	return client
}

func NewRedisProcessor(redisClient *redis.Client, db *sql.DB) *RedisProcessor {
	processor := &RedisProcessor{
		redisClient: redisClient,
		db:          db,
		ctx:         context.Background(),
	}

	// Auto-start the processor
	go processor.processEvents()

	log.Println("Redis processor initialized")
	return processor
}

func (rp *RedisProcessor) processEvents() {
	log.Println("Starting events processor...")

	// Create consumer group if it doesn't exist
	err := rp.redisClient.XGroupCreate(rp.ctx, "events:live", "analytics-processor", "0").Err()
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		log.Printf("Error creating consumer group: %v", err)
	}

	for {
		// Read from consumer group - only gets unprocessed messages
		result, err := rp.redisClient.XReadGroup(rp.ctx, &redis.XReadGroupArgs{
			Group:    "analytics-processor",
			Consumer: "worker-1",
			Streams:  []string{"events:live", ">"}, // ">" means only new messages
			Block:    0,
		}).Result()

		if err != nil {
			log.Printf("Error reading from stream: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		for _, stream := range result {
			for _, message := range stream.Messages {
				// Parse the event
				eventJSON := message.Values["event"].(string)
				var event types.RedisEvent
				err := json.Unmarshal([]byte(eventJSON), &event)
				if err != nil {
					log.Printf("Error unmarshaling event: %v", err)
					// Acknowledge even failed messages to prevent infinite retries
					rp.redisClient.XAck(rp.ctx, "events:live", "analytics-processor", message.ID)
					continue
				}

				// Marshal the data field to JSON for database storage
				jsonData, err := json.Marshal(event.EventData)
				if err != nil {
					log.Printf("Error marshaling event data: %v", err)
					rp.redisClient.XAck(rp.ctx, "events:live", "analytics-processor", message.ID)
					continue
				}

				// Store in PostgreSQL
				_, err = rp.db.Exec(`
					INSERT INTO events (user_id, event_type, client_id, event_url, event_data)
					VALUES ($1, $2, $3, $4, $5)`,
					event.UserID, event.EventType, event.ClientId, event.EventUrl, jsonData)

				if err != nil {
					log.Printf("DB insert error: %v", err)
					// Don't acknowledge failed DB inserts - they'll be retried
					continue
				}

				// Acknowledge successful processing - removes from pending list
				rp.redisClient.XAck(rp.ctx, "events:live", "analytics-processor", message.ID)
				log.Printf("Processed %s event for user: %s", event.EventType, event.UserID)
			}
		}
	}
}
