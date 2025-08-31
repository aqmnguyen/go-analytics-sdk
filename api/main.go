package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"analytics-api/config"
	"analytics-api/handlers"
	"analytics-api/workers"

	"github.com/redis/go-redis/v9"
)

func main() {

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	redisClient := ConnectRedis()
	if redisClient == nil {
		log.Fatal("Failed to connect to Redis")
	}
	defer redisClient.Close()

	workers.NewRedisProcessor(redisClient, db)

	http.HandleFunc("/event/click", handlers.ClickHandler(redisClient))
	http.HandleFunc("/event/pageview", handlers.PageviewHandler(redisClient))

	fmt.Fprintf(os.Stdout, "Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
