package config

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(cfg *Config) (*sql.DB, error) {
	var db *sql.DB
	var err error

	if cfg.Environment == "prod" {
		// For production, we'll use SQLite for now
		// In a real deployment, you would configure Turso properly
		log.Println("Warning: Production mode using SQLite. Configure Turso for production.")
		dbPath := "../local/db.db"
		db, err = sql.Open("sqlite3", dbPath)
	} else {
		// For SQLite connection (remove sqlite:// prefix)
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
