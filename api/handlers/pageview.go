package handlers

import (
	"analytics-api/types"
	"analytics-api/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func PageviewHandler(db *sql.DB) http.HandlerFunc {

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

		dbErr := InsertPageviewEvent(db, pageviewEvent)
		if dbErr != nil {
			log.Printf("Error inserting pageview event: %v", dbErr)
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

func InsertPageviewEvent(db *sql.DB, pageviewEvent types.PageviewEvent) error {
	eventData := map[string]interface{}{
		"referrer":   getStringValue(pageviewEvent.Referrer),
		"ip_address": getStringValue(pageviewEvent.IpAddress),
		"user_agent": getStringValue(pageviewEvent.UserAgent),
	}

	jsonData, err := json.Marshal(eventData)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO events (user_id, event_type, event_url, event_data) VALUES ($1, $2, $3, $4)",
		pageviewEvent.UserId, pageviewEvent.Event, pageviewEvent.Url, jsonData)
	return err
}

func getStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
