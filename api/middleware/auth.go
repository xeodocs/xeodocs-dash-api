package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xeodocs/xeodocs-dash-api/internal/service"
)

type AuthMiddleware struct {
	userService *service.UserService
}

func NewAuthMiddleware(userService *service.UserService) *AuthMiddleware {
	return &AuthMiddleware{userService: userService}
}

func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>" format
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format. Use 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		token := tokenParts[1]

		// Validate session token
		user, err := m.userService.ValidateSession(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired session token",
			})
			c.Abort()
			return
		}

		// Store user in context for use in handlers
		c.Set("user", user)
		c.Set("user_id", user.ID)

		c.Next()
	}
}

func (m *AuthMiddleware) OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		// Extract token from "Bearer <token>" format
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.Next()
			return
		}

		token := tokenParts[1]

		// Validate session token
		user, err := m.userService.ValidateSession(token)
		if err != nil {
			c.Next()
			return
		}

		// Store user in context for use in handlers
		c.Set("user", user)
		c.Set("user_id", user.ID)

		c.Next()
	}
}
