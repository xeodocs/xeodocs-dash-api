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

// GetWebsites godoc
// @Summary Get all websites
// @Description Get list of all websites
// @Tags Websites
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string][]models.Website "List of websites"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/websites [get]
func (h *WebsiteHandler) GetWebsites(c *gin.Context) {
	websites, err := h.websiteService.GetAllWebsites()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"websites": websites})
}

// GetWebsite godoc
// @Summary Get website by ID
// @Description Get a specific website by its ID
// @Tags Websites
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Website ID"
// @Success 200 {object} map[string]models.Website "Website details"
// @Failure 400 {object} map[string]string "Invalid website ID"
// @Failure 404 {object} map[string]string "Website not found"
// @Router /api/v1/websites/{id} [get]
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

// GetWebsiteBySlug godoc
// @Summary Get website by slug
// @Description Get a specific website by its slug
// @Tags Websites
// @Accept json
// @Produce json
// @Security Bearer
// @Param slug path string true "Website slug"
// @Success 200 {object} map[string]models.Website "Website details"
// @Failure 404 {object} map[string]string "Website not found"
// @Router /api/v1/websites/slug/{slug} [get]
func (h *WebsiteHandler) GetWebsiteBySlug(c *gin.Context) {
	slug := c.Param("slug")

	website, err := h.websiteService.GetWebsiteBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"website": website})
}

// CreateWebsite godoc
// @Summary Create new website
// @Description Create a new website
// @Tags Websites
// @Accept json
// @Produce json
// @Security Bearer
// @Param website body models.CreateWebsiteRequest true "Website creation data"
// @Success 201 {object} map[string]models.Website "Website created successfully"
// @Failure 400 {object} map[string]string "Bad request or validation error"
// @Router /api/v1/websites [post]
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

// UpdateWebsite godoc
// @Summary Update website
// @Description Update website information
// @Tags Websites
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Website ID"
// @Param website body models.UpdateWebsiteRequest true "Website update data"
// @Success 200 {object} map[string]models.Website "Website updated successfully"
// @Failure 400 {object} map[string]string "Bad request or validation error"
// @Failure 404 {object} map[string]string "Website not found"
// @Router /api/v1/websites/{id} [put]
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

// DeleteWebsite godoc
// @Summary Delete website
// @Description Delete a website
// @Tags Websites
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Website ID"
// @Success 200 {object} map[string]string "Website deleted successfully"
// @Failure 400 {object} map[string]string "Invalid website ID"
// @Failure 404 {object} map[string]string "Website not found"
// @Router /api/v1/websites/{id} [delete]
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
