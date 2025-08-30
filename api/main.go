package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"analytics-api/config"
	"analytics-api/handlers"
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

	http.HandleFunc("/event/click", handlers.ClickHandler(db))
	http.HandleFunc("/event/pageview", handlers.PageviewHandler(db))

	fmt.Fprintf(os.Stdout, "Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
