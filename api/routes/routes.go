package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/xeodocs/xeodocs-dash-api/api/handlers"
	"github.com/xeodocs/xeodocs-dash-api/api/middleware"
	"github.com/xeodocs/xeodocs-dash-api/internal/repository"
	"github.com/xeodocs/xeodocs-dash-api/internal/service"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	websiteRepo := repository.NewWebsiteRepository(db)
	pageRepo := repository.NewPageRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	websiteService := service.NewWebsiteService(websiteRepo)
	pageService := service.NewPageService(pageRepo, websiteRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	websiteHandler := handlers.NewWebsiteHandler(websiteService)
	pageHandler := handlers.NewPageHandler(pageService)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(userService)

	// Setup Gin router
	r := gin.Default()

	// Global middleware
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	// Health check endpoint (no auth required)
	// HealthCheck godoc
	// @Summary Health check
	// @Description Check if the API service is running
	// @Tags Health
	// @Accept json
	// @Produce json
	// @Success 200 {object} map[string]string "Service is healthy"
	// @Router /health [get]
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	// Swagger documentation endpoints
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Authentication routes (no auth required)
		auth := v1.Group("/auth")
		{
			auth.POST("/login", userHandler.Login)
			auth.POST("/logout", userHandler.Logout)
			auth.GET("/me", authMiddleware.RequireAuth(), userHandler.GetCurrentUser)
		}

		// Protected routes (require authentication)
		protected := v1.Group("/")
		protected.Use(authMiddleware.RequireAuth())
		{
			// User routes
			users := protected.Group("/users")
			{
				users.GET("", userHandler.GetUsers)
				users.GET("/:id", userHandler.GetUser)
				users.POST("", userHandler.CreateUser)
				users.PUT("/:id", userHandler.UpdateUser)
				users.DELETE("/:id", userHandler.DeleteUser)
			}

			// Website routes
			websites := protected.Group("/websites")
			{
				websites.GET("", websiteHandler.GetWebsites)
				websites.GET("/:id", websiteHandler.GetWebsite)
				websites.GET("/slug/:slug", websiteHandler.GetWebsiteBySlug)
				websites.POST("", websiteHandler.CreateWebsite)
				websites.PUT("/:id", websiteHandler.UpdateWebsite)
				websites.DELETE("/:id", websiteHandler.DeleteWebsite)
			}

			// Page routes
			pages := protected.Group("/pages")
			{
				pages.GET("", pageHandler.GetPages) // Supports ?website_id=X query param
				pages.GET("/:id", pageHandler.GetPage)
				pages.GET("/slug/:slug", pageHandler.GetPageBySlug)
				pages.POST("", pageHandler.CreatePage)
				pages.PUT("/:id", pageHandler.UpdatePage)
				pages.DELETE("/:id", pageHandler.DeletePage)
			}
		}
	}

	return r
}
