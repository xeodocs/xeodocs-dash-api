package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Error:   message,
	})
}

func ValidationErrorResponse(c *gin.Context, err error) {
	ErrorResponse(c, http.StatusBadRequest, err.Error())
}

func NotFoundResponse(c *gin.Context, resource string) {
	ErrorResponse(c, http.StatusNotFound, resource+" not found")
}

func InternalErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message)
}

func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message)
}
