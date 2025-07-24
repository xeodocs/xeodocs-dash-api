package service

import (
	"fmt"
	"time"

	"github.com/xeodocs/xeodocs-dash-api/internal/models"
	"github.com/xeodocs/xeodocs-dash-api/internal/repository"
)

type PageService struct {
	pageRepo    *repository.PageRepository
	websiteRepo *repository.WebsiteRepository
}

func NewPageService(pageRepo *repository.PageRepository, websiteRepo *repository.WebsiteRepository) *PageService {
	return &PageService{
		pageRepo:    pageRepo,
		websiteRepo: websiteRepo,
	}
}

func (s *PageService) CreatePage(req *models.CreatePageRequest) (*models.Page, error) {
	// Verify website exists
	_, err := s.websiteRepo.GetByID(req.WebsiteID)
	if err != nil {
		return nil, fmt.Errorf("website not found: %w", err)
	}

	// Check if page with same slug already exists
	existingPage, _ := s.pageRepo.GetBySlug(req.Slug)
	if existingPage != nil {
		return nil, fmt.Errorf("page with slug %s already exists", req.Slug)
	}

	// Set default tags if empty
	tags := req.Tags
	if tags == "" {
		tags = "[]"
	}

	page := &models.Page{
		WebsiteID:          req.WebsiteID,
		Title:              req.Title,
		Slug:               req.Slug,
		Description:        req.Description,
		MarkdownContent:    req.MarkdownContent,
		Tags:               tags,
		FreezeStatus:       req.FreezeStatus,
		Status:             req.Status,
		ScheduledPublishAt: req.ScheduledPublishAt,
	}

	err = s.pageRepo.Create(page)
	if err != nil {
		return nil, fmt.Errorf("failed to create page: %w", err)
	}

	return page, nil
}

func (s *PageService) GetPageByID(id int) (*models.Page, error) {
	return s.pageRepo.GetByID(id)
}

func (s *PageService) GetPageBySlug(slug string) (*models.Page, error) {
	return s.pageRepo.GetBySlug(slug)
}

func (s *PageService) GetAllPages() ([]*models.Page, error) {
	return s.pageRepo.GetAll()
}

func (s *PageService) GetPagesByWebsiteID(websiteID int) ([]*models.Page, error) {
	// Verify website exists
	_, err := s.websiteRepo.GetByID(websiteID)
	if err != nil {
		return nil, fmt.Errorf("website not found: %w", err)
	}

	return s.pageRepo.GetByWebsiteID(websiteID)
}

func (s *PageService) UpdatePage(id int, req *models.UpdatePageRequest) (*models.Page, error) {
	page, err := s.pageRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Track if status is changing for last_status_change_at
	statusChanged := false

	// Update fields if provided
	if req.Title != "" {
		page.Title = req.Title
	}
	if req.Slug != "" {
		// Check if new slug is already taken by another page
		existingPage, _ := s.pageRepo.GetBySlug(req.Slug)
		if existingPage != nil && existingPage.ID != id {
			return nil, fmt.Errorf("slug %s is already taken", req.Slug)
		}
		page.Slug = req.Slug
	}
	if req.Description != "" {
		page.Description = req.Description
	}
	if req.MarkdownContent != "" {
		page.MarkdownContent = req.MarkdownContent
	}
	if req.Tags != "" {
		page.Tags = req.Tags
	}
	if req.FreezeStatus != nil {
		page.FreezeStatus = *req.FreezeStatus
	}
	if req.Status != "" && req.Status != page.Status {
		page.Status = req.Status
		statusChanged = true
	}
	if req.ScheduledPublishAt != nil {
		page.ScheduledPublishAt = req.ScheduledPublishAt
	}

	// Update last_status_change_at if status changed
	if statusChanged {
		page.LastStatusChangeAt = time.Now()
	}

	err = s.pageRepo.Update(id, page)
	if err != nil {
		return nil, err
	}

	return page, nil
}

func (s *PageService) DeletePage(id int) error {
	return s.pageRepo.Delete(id)
}
