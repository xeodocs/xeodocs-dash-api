package config

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func InitDatabase(cfg *Config) (*sql.DB, error) {
	var db *sql.DB
	var err error

	if cfg.Environment == "prod" {
		// For production, use Turso
		log.Println("Connecting to Turso database...")
		db, err = sql.Open("libsql", fmt.Sprintf("%s?authToken=%s", cfg.DatabaseURL, cfg.TursoAuthToken))
	} else {
		// For development, use SQLite
		log.Println("Connecting to SQLite database...")
		dbPath := strings.TrimPrefix(cfg.DatabaseURL, "sqlite://")
		db, err = sql.Open("sqlite3", dbPath)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Printf("Successfully connected to database (%s)", cfg.Environment)
	return db, nil
}
