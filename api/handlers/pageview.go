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

func PageviewHandler(redisClient *redis.Client) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var pageviewEvent types.PageviewEvent
		err := json.NewDecoder(r.Body).Decode(&pageviewEvent)
		if err != nil {
			log.Printf("Error decoding JSON: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Validate required fields using tags
		if missing := utils.ValidateRequired(pageviewEvent); len(missing) > 0 {
			log.Printf("Missing required fields: %v", missing)
			http.Error(w, "Missing required fields: "+strings.Join(missing, ", "), http.StatusBadRequest)
			return
		}

		log.Printf("Pageview event: %v", pageviewEvent)

		// Publish to Redis stream
		redisErr := PublishPageviewEvent(redisClient, pageviewEvent)
		if redisErr != nil {
			log.Printf("Error publishing click event to Redis: %v", redisErr)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Printf("Pageview event inserted successfully")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		jsonBytes, err := json.Marshal(pageviewEvent)
		if err != nil {
			log.Printf("Error marshalling JSON response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)
	}
}

func PublishPageviewEvent(redisClient *redis.Client, pageviewEvent types.PageviewEvent) error {

	// Create Event object for Redis stream
	event := map[string]interface{}{
		"event_type": "pageview",
		"client_id":  pageviewEvent.ClientId,
		"user_id":    pageviewEvent.UserId,
		"url":        pageviewEvent.Url,
		"event_data": map[string]interface{}{
			"referrer":   utils.GetStringValue(pageviewEvent.Referrer),
			"ip_address": utils.GetStringValue(pageviewEvent.IpAddress),
			"user_agent": utils.GetStringValue(pageviewEvent.UserAgent),
		},
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
