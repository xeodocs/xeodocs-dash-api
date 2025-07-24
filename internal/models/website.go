package models

import (
	"time"
)

type Website struct {
	ID             int       `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Slug           string    `json:"slug" db:"slug"`
	Description    string    `json:"description" db:"description"`
	Slogan         string    `json:"slogan" db:"slogan"`
	Domain         string    `json:"domain" db:"domain"`
	GitRepoOwner   string    `json:"git_repo_owner" db:"git_repo_owner"`
	GitRepoName    string    `json:"git_repo_name" db:"git_repo_name"`
	GitRepoBranch  string    `json:"git_repo_branch" db:"git_repo_branch"`
	GitAPIToken    string    `json:"-" db:"git_api_token"`
	Config         string    `json:"config" db:"config"`
	LanguageCode   string    `json:"language_code" db:"language_code"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Request/Response DTOs
type CreateWebsiteRequest struct {
	Name          string `json:"name" binding:"required"`
	Slug          string `json:"slug" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Slogan        string `json:"slogan" binding:"required"`
	Domain        string `json:"domain" binding:"required"`
	GitRepoOwner  string `json:"git_repo_owner" binding:"required"`
	GitRepoName   string `json:"git_repo_name" binding:"required"`
	GitRepoBranch string `json:"git_repo_branch" binding:"required"`
	GitAPIToken   string `json:"git_api_token" binding:"required"`
	Config        string `json:"config" binding:"required"`
	LanguageCode  string `json:"language_code" binding:"required,len=2"`
}

type UpdateWebsiteRequest struct {
	Name          string `json:"name" binding:"omitempty"`
	Slug          string `json:"slug" binding:"omitempty"`
	Description   string `json:"description" binding:"omitempty"`
	Slogan        string `json:"slogan" binding:"omitempty"`
	Domain        string `json:"domain" binding:"omitempty"`
	GitRepoOwner  string `json:"git_repo_owner" binding:"omitempty"`
	GitRepoName   string `json:"git_repo_name" binding:"omitempty"`
	GitRepoBranch string `json:"git_repo_branch" binding:"omitempty"`
	GitAPIToken   string `json:"git_api_token" binding:"omitempty"`
	Config        string `json:"config" binding:"omitempty"`
	LanguageCode  string `json:"language_code" binding:"omitempty,len=2"`
}
