package db

import (
	"auth-server/read-mail/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDBConnection() (*sql.DB, error) {
	// Load application configuration
	appConfig := config.Config()
	// Replace the connection parameters with your PostgreSQL database details
	db, err := sql.Open("postgres", "postgres://"+appConfig.UserName+":"+appConfig.Password+"@"+appConfig.DatabaseURL+"/"+appConfig.DatabaseName+"?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
