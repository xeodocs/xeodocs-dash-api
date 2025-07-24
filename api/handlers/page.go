package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xeodocs/xeodocs-dash-api/internal/models"
	"github.com/xeodocs/xeodocs-dash-api/internal/service"
)

type PageHandler struct {
	pageService *service.PageService
}

func NewPageHandler(pageService *service.PageService) *PageHandler {
	return &PageHandler{pageService: pageService}
}

// GET /pages
func (h *PageHandler) GetPages(c *gin.Context) {
	// Check if filtering by website_id
	websiteIDStr := c.Query("website_id")
	if websiteIDStr != "" {
		websiteID, err := strconv.Atoi(websiteIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid website_id parameter"})
			return
		}

		pages, err := h.pageService.GetPagesByWebsiteID(websiteID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"pages": pages})
		return
	}

	pages, err := h.pageService.GetAllPages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pages": pages})
}

// GET /pages/:id
func (h *PageHandler) GetPage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	page, err := h.pageService.GetPageByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"page": page})
}

// GET /pages/slug/:slug
func (h *PageHandler) GetPageBySlug(c *gin.Context) {
	slug := c.Param("slug")

	page, err := h.pageService.GetPageBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"page": page})
}

// POST /pages
func (h *PageHandler) CreatePage(c *gin.Context) {
	var req models.CreatePageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, err := h.pageService.CreatePage(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"page": page})
}

// PUT /pages/:id
func (h *PageHandler) UpdatePage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	var req models.UpdatePageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, err := h.pageService.UpdatePage(id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"page": page})
}

// DELETE /pages/:id
func (h *PageHandler) DeletePage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	err = h.pageService.DeletePage(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page deleted successfully"})
}
