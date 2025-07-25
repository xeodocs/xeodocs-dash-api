package config

import (
	"os"
	"strings"
)

type Config struct {
	Environment    string
	DatabaseURL    string
	TursoAuthToken string
	Port           string
}

func Load() *Config {
	env := getEnv("ENVIRONMENT", "dev")

	var databaseURL string
	if env == "prod" {
		databaseURL = getEnv("TURSO_DB_URL", "")
		if databaseURL != "" && !strings.HasPrefix(databaseURL, "libsql://") && !strings.HasPrefix(databaseURL, "file://") && !strings.HasPrefix(databaseURL, "https://") && !strings.HasPrefix(databaseURL, "http://") && !strings.HasPrefix(databaseURL, "wss://") && !strings.HasPrefix(databaseURL, "ws://") {
			databaseURL = "libsql://" + databaseURL
		}
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
