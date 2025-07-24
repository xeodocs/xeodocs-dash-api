package service

import (
	"fmt"

	"github.com/xeodocs/xeodocs-dash-api/internal/models"
	"github.com/xeodocs/xeodocs-dash-api/internal/repository"
)

type WebsiteService struct {
	websiteRepo *repository.WebsiteRepository
}

func NewWebsiteService(websiteRepo *repository.WebsiteRepository) *WebsiteService {
	return &WebsiteService{websiteRepo: websiteRepo}
}

func (s *WebsiteService) CreateWebsite(req *models.CreateWebsiteRequest) (*models.Website, error) {
	// Check if website with same name or slug already exists
	existingBySlug, _ := s.websiteRepo.GetBySlug(req.Slug)
	if existingBySlug != nil {
		return nil, fmt.Errorf("website with slug %s already exists", req.Slug)
	}

	website := &models.Website{
		Name:          req.Name,
		Slug:          req.Slug,
		Description:   req.Description,
		Slogan:        req.Slogan,
		Domain:        req.Domain,
		GitRepoOwner:  req.GitRepoOwner,
		GitRepoName:   req.GitRepoName,
		GitRepoBranch: req.GitRepoBranch,
		GitAPIToken:   req.GitAPIToken,
		Config:        req.Config,
		LanguageCode:  req.LanguageCode,
	}

	err := s.websiteRepo.Create(website)
	if err != nil {
		return nil, fmt.Errorf("failed to create website: %w", err)
	}

	return website, nil
}

func (s *WebsiteService) GetWebsiteByID(id int) (*models.Website, error) {
	return s.websiteRepo.GetByID(id)
}

func (s *WebsiteService) GetWebsiteBySlug(slug string) (*models.Website, error) {
	return s.websiteRepo.GetBySlug(slug)
}

func (s *WebsiteService) GetAllWebsites() ([]*models.Website, error) {
	return s.websiteRepo.GetAll()
}

func (s *WebsiteService) UpdateWebsite(id int, req *models.UpdateWebsiteRequest) (*models.Website, error) {
	website, err := s.websiteRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Name != "" {
		website.Name = req.Name
	}
	if req.Slug != "" {
		// Check if new slug is already taken by another website
		existingWebsite, _ := s.websiteRepo.GetBySlug(req.Slug)
		if existingWebsite != nil && existingWebsite.ID != id {
			return nil, fmt.Errorf("slug %s is already taken", req.Slug)
		}
		website.Slug = req.Slug
	}
	if req.Description != "" {
		website.Description = req.Description
	}
	if req.Slogan != "" {
		website.Slogan = req.Slogan
	}
	if req.Domain != "" {
		website.Domain = req.Domain
	}
	if req.GitRepoOwner != "" {
		website.GitRepoOwner = req.GitRepoOwner
	}
	if req.GitRepoName != "" {
		website.GitRepoName = req.GitRepoName
	}
	if req.GitRepoBranch != "" {
		website.GitRepoBranch = req.GitRepoBranch
	}
	if req.GitAPIToken != "" {
		website.GitAPIToken = req.GitAPIToken
	}
	if req.Config != "" {
		website.Config = req.Config
	}
	if req.LanguageCode != "" {
		website.LanguageCode = req.LanguageCode
	}

	err = s.websiteRepo.Update(id, website)
	if err != nil {
		return nil, err
	}

	return website, nil
}

func (s *WebsiteService) DeleteWebsite(id int) error {
	return s.websiteRepo.Delete(id)
}
