package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const GRAPH_URL = "https://graph.microsoft.com/v1.0"

type AppConfig struct {
	DatabaseURL  string
	DatabaseName string
	TableName    string
	UserName     string
	Password     string
}

func Config() *AppConfig {

	return &AppConfig{
		DatabaseURL:  getEnv("DATABASE_URL", ""),
		DatabaseName: getEnv("DATABASE_NAME", ""),
		TableName:    getEnv("DATABASE_TABLE", ""),
		UserName:     getEnv("USERNAME", ""),
		Password:     getEnv("PASSWORD", ""),
	}
}

func getEnv(key, defaultValue string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
