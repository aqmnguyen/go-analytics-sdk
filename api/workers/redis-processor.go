package workers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisProcessor struct {
	redisClient *redis.Client
	db          *sql.DB
	ctx         context.Context
}

type Event struct {
	EventType string                 `json:"event_type"`
	UserID    string                 `json:"user_id"`
	URL       string                 `json:"url"`
	EventData map[string]interface{} `json:"event_data"`
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

	// Track last processed ID to only read new events
	lastID := "0"

	for {
		// Read from Redis stream starting from last processed ID
		result, err := rp.redisClient.XRead(rp.ctx, &redis.XReadArgs{
			Streams: []string{"events:live", lastID},
			Block:   0,
		}).Result()

		if err != nil {
			log.Printf("Error reading from stream: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		for _, stream := range result {
			for _, message := range stream.Messages {
				// Update last processed ID to avoid reprocessing
				lastID = message.ID

				// Parse the event
				eventJSON := message.Values["event"].(string)
				var event Event
				err := json.Unmarshal([]byte(eventJSON), &event)
				if err != nil {
					log.Printf("Error unmarshaling event: %v", err)
					continue
				}

				// Marshal the data field to JSON for database storage
				jsonData, err := json.Marshal(event.EventData)
				if err != nil {
					log.Printf("Error marshaling event data: %v", err)
					continue
				}

				// Store in PostgreSQL
				_, err = rp.db.Exec(`
					INSERT INTO events (user_id, event_type, event_url, event_data)
					VALUES ($1, $2, $3, $4)`,
					event.UserID, event.EventType, event.URL, jsonData)

				if err != nil {
					log.Printf("DB insert error: %v", err)
					continue
				}

				log.Printf("Processed %s event for user: %s", event.EventType, event.UserID)
			}
		}
	}
}
