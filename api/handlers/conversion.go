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

func ConversionHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var conversionEvent types.ConversionEvent
		err := json.NewDecoder(r.Body).Decode(&conversionEvent)
		if err != nil {
			log.Printf("Error decoding JSON: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Validate required fields using tags
		if missing := utils.ValidateRequired(conversionEvent); len(missing) > 0 {
			log.Printf("Missing required fields: %v", missing)
			http.Error(w, "Missing required fields: "+strings.Join(missing, ", "), http.StatusBadRequest)
			return
		}

		log.Printf("Conversion event: %v", conversionEvent)

		// Publish to Redis stream
		redisErr := PublishConversionEvent(redisClient, conversionEvent)
		if redisErr != nil {
			log.Printf("Error publishing conversion event to Redis: %v", redisErr)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Printf("Conversion event published to Redis successfully")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		jsonBytes, err := json.Marshal(conversionEvent)
		if err != nil {
			log.Printf("Error marshalling JSON response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)

	}
}

func PublishConversionEvent(redisClient *redis.Client, conversionEvent types.ConversionEvent) error {

	// Create Event object for Redis stream
	event := map[string]interface{}{
		"event_type": "conversion",
		"client_id":  conversionEvent.ClientId,
		"user_id":    conversionEvent.UserId,
		"event_url":  conversionEvent.EventUrl,
		"event_data": conversionEvent.EventData,
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
