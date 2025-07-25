package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment    string
	DatabaseURL    string
	TursoAuthToken string
	Port           string
}

func Load() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	env := getEnv("ENVIRONMENT", "dev")

	var databaseURL string
	if env == "prod" {
		databaseURL = getEnv("TURSO_DB_URL", "")
	} else {
		databaseURL = "sqlite://./local/db.db"
	}

	return &Config{
		Environment:    env,
		DatabaseURL:    databaseURL,
		TursoAuthToken: getEnv("TURSO_AUTH_TOKEN", ""),
		Port:           getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
