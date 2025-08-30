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

func ClickHandler(db *sql.DB) http.HandlerFunc {
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

		log.Printf("Click event: %v", clickEvent)

		dbErr := InsertClickEvent(db, clickEvent)
		if dbErr != nil {
			log.Printf("Error inserting click event: %v", dbErr)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Printf("Click event inserted successfully")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		jsonBytes, err := json.Marshal(clickEvent)
		if err != nil {
			log.Printf("Error marshalling JSON response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)

	}
}

func InsertClickEvent(db *sql.DB, clickEvent types.ClickEvent) error {
	eventData := map[string]interface{}{
		"element":    clickEvent.Element,
		"referrer":   getStringValue(clickEvent.Referrer),
		"ip_address": getStringValue(clickEvent.IpAddress),
		"user_agent": getStringValue(clickEvent.UserAgent),
	}

	jsonData, err := json.Marshal(eventData)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO events (user_id, event_type, event_url, event_data) VALUES ($1, $2, $3, $4)", clickEvent.UserId, clickEvent.Event, clickEvent.Url, jsonData)
	return err
}
