package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xeodocs/xeodocs-dash-api/internal/models"
	"github.com/xeodocs/xeodocs-dash-api/internal/service"
)

type WebsiteHandler struct {
	websiteService *service.WebsiteService
}

func NewWebsiteHandler(websiteService *service.WebsiteService) *WebsiteHandler {
	return &WebsiteHandler{websiteService: websiteService}
}

// GET /websites
func (h *WebsiteHandler) GetWebsites(c *gin.Context) {
	websites, err := h.websiteService.GetAllWebsites()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"websites": websites})
}

// GET /websites/:id
func (h *WebsiteHandler) GetWebsite(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid website ID"})
		return
	}

	website, err := h.websiteService.GetWebsiteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"website": website})
}

// GET /websites/slug/:slug
func (h *WebsiteHandler) GetWebsiteBySlug(c *gin.Context) {
	slug := c.Param("slug")

	website, err := h.websiteService.GetWebsiteBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"website": website})
}

// POST /websites
func (h *WebsiteHandler) CreateWebsite(c *gin.Context) {
	var req models.CreateWebsiteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	website, err := h.websiteService.CreateWebsite(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"website": website})
}

// PUT /websites/:id
func (h *WebsiteHandler) UpdateWebsite(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid website ID"})
		return
	}

	var req models.UpdateWebsiteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	website, err := h.websiteService.UpdateWebsite(id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"website": website})
}

// DELETE /websites/:id
func (h *WebsiteHandler) DeleteWebsite(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid website ID"})
		return
	}

	err = h.websiteService.DeleteWebsite(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Website deleted successfully"})
}
