package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"analytics-api/config"
	"analytics-api/handlers"
	"analytics-api/workers"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	redisClient := workers.ConnectRedis()
	if redisClient == nil {
		log.Fatal("Failed to connect to Redis")
	}
	defer redisClient.Close()

	workers.NewRedisProcessor(redisClient, db)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/event/click", enableCORS(handlers.ClickHandler(redisClient)))
	http.HandleFunc("/event/pageview", enableCORS(handlers.PageviewHandler(redisClient)))
	http.HandleFunc("/event/conversion", enableCORS(handlers.ConversionHandler(redisClient)))

	fmt.Fprintf(os.Stdout, "Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// CORS middleware function
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next(w, r)
	}
}
