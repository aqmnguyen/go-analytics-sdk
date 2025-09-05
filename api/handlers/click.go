package handlers

import (
	"analytics-api/types"
	"analytics-api/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/redis/go-redis/v9"
)

func ClickHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var clickEvent types.ClickEvent
		err := json.NewDecoder(r.Body).Decode(&clickEvent)
		if err != nil {
			log.Printf("Error decoding JSON: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Validate required fields using tags
		if missing := utils.ValidateRequired(clickEvent); len(missing) > 0 {
			log.Printf("Missing required fields: %v", missing)
			http.Error(w, "Missing required fields: "+strings.Join(missing, ", "), http.StatusBadRequest)
			return
		}

		log.Printf("Click event received: %v", clickEvent)

		// Publish to Redis stream
		redisErr := PublishClickEvent(redisClient, clickEvent)
		if redisErr != nil {
			log.Printf("Error publishing click event to Redis: %v", redisErr)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Printf("Click event published to Redis successfully")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		jsonBytes, err := json.Marshal(clickEvent)
		if err != nil {
			log.Printf("Error marshalling JSON response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)

	}
}

func PublishClickEvent(redisClient *redis.Client, clickEvent types.ClickEvent) error {

	// Create Event object for Redis stream
	event := map[string]interface{}{
		"event_type": "click",
		"client_id":  clickEvent.ClientId,
		"user_id":    clickEvent.UserId,
		"event_url":  clickEvent.EventUrl,
		"event_data": clickEvent.EventData,
	}

	// Marshal the event to JSON
	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Publish to Redis stream
	ctx := context.Background()
	err = redisClient.XAdd(ctx, &redis.XAddArgs{
		Stream: "events:live",
		Values: map[string]interface{}{
			"event": string(jsonData),
		},
	}).Err()

	return err
}
