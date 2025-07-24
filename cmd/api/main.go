package main

import (
	"log"

	"github.com/xeodocs/xeodocs-dash-api/api/routes"
	"github.com/xeodocs/xeodocs-dash-api/config"
)

func main() {
	// Load configuration
	cfg := config.Load()
	log.Printf("Starting server in %s mode", cfg.Environment)

	// Initialize database connection
	db, err := config.InitDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Setup routes
	router := routes.SetupRoutes(db)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
