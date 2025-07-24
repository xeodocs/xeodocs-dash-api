package config

import (
	"os"
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
		databaseURL = "libsql://xeodocs-db-xeodocs.aws-us-east-1.turso.io"
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
