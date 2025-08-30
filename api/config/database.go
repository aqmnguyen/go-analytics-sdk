package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// dbURL := os.Getenv("DATABASE_URL")
	dbURL := "postgresql://analytics_user:analytics_password@localhost:5433/analytics?sslmode=disable"
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
