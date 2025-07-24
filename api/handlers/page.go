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

// GetPages godoc
// @Summary Get all pages
// @Description Get list of all pages, optionally filtered by website_id
// @Tags Pages
// @Accept json
// @Produce json
// @Security Bearer
// @Param website_id query int false "Filter by website ID"
// @Success 200 {object} map[string][]models.Page "List of pages"
// @Failure 400 {object} map[string]string "Invalid website_id parameter"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/pages [get]
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

// GetPage godoc
// @Summary Get page by ID
// @Description Get a specific page by its ID
// @Tags Pages
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Page ID"
// @Success 200 {object} map[string]models.Page "Page details"
// @Failure 400 {object} map[string]string "Invalid page ID"
// @Failure 404 {object} map[string]string "Page not found"
// @Router /api/v1/pages/{id} [get]
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

// GetPageBySlug godoc
// @Summary Get page by slug
// @Description Get a specific page by its slug
// @Tags Pages
// @Accept json
// @Produce json
// @Security Bearer
// @Param slug path string true "Page slug"
// @Success 200 {object} map[string]models.Page "Page details"
// @Failure 404 {object} map[string]string "Page not found"
// @Router /api/v1/pages/slug/{slug} [get]
func (h *PageHandler) GetPageBySlug(c *gin.Context) {
	slug := c.Param("slug")

	page, err := h.pageService.GetPageBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"page": page})
}

// CreatePage godoc
// @Summary Create new page
// @Description Create a new page
// @Tags Pages
// @Accept json
// @Produce json
// @Security Bearer
// @Param page body models.CreatePageRequest true "Page creation data"
// @Success 201 {object} map[string]models.Page "Page created successfully"
// @Failure 400 {object} map[string]string "Bad request or validation error"
// @Router /api/v1/pages [post]
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

// UpdatePage godoc
// @Summary Update page
// @Description Update page information
// @Tags Pages
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Page ID"
// @Param page body models.UpdatePageRequest true "Page update data"
// @Success 200 {object} map[string]models.Page "Page updated successfully"
// @Failure 400 {object} map[string]string "Bad request or validation error"
// @Failure 404 {object} map[string]string "Page not found"
// @Router /api/v1/pages/{id} [put]
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

// DeletePage godoc
// @Summary Delete page
// @Description Delete a page
// @Tags Pages
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Page ID"
// @Success 200 {object} map[string]string "Page deleted successfully"
// @Failure 400 {object} map[string]string "Invalid page ID"
// @Failure 404 {object} map[string]string "Page not found"
// @Router /api/v1/pages/{id} [delete]
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
