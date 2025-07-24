// Package main XeoDocs Dash API
//
// RESTful API service for managing users, websites, and pages with session-based authentication.
//
// Terms Of Service: http://swagger.io/terms/
//
// Schemes: http
// Host: localhost:8080
// BasePath: /
// Version: 1.0.0
// Contact: XeoDocs Team<support@xeodocs.com> http://www.xeodocs.com
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Security:
// - Bearer: []
//
// SecurityDefinitions:
// Bearer:
//   type: apiKey
//   name: Authorization
//   in: header
//   description: Bearer token for authentication. Format: 'Bearer {token}'
//
// @title XeoDocs Dash API
// @version 1.0
// @description RESTful API service for managing users, websites, and pages with session-based authentication.
// @termsOfService http://swagger.io/terms/
// @contact.name XeoDocs Team
// @contact.email support@xeodocs.com
// @contact.url http://www.xeodocs.com
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Bearer token for authentication. Format: 'Bearer {token}'
package main

import (
	"log"

	"github.com/xeodocs/xeodocs-dash-api/api/routes"
	"github.com/xeodocs/xeodocs-dash-api/config"
	_ "github.com/xeodocs/xeodocs-dash-api/docs" // Import generated docs
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
