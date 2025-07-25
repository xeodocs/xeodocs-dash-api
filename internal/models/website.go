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
	GitRepoOwner   string    `json:"gitRepoOwner" db:"git_repo_owner"`
	GitRepoName    string    `json:"gitRepoName" db:"git_repo_name"`
	GitRepoBranch  string    `json:"gitRepoBranch" db:"git_repo_branch"`
	GitAPIToken    string    `json:"-" db:"git_api_token"`
	Config         string    `json:"config" db:"config"`
	LanguageCode   string    `json:"languageCode" db:"language_code"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`
}

// Request/Response DTOs
type CreateWebsiteRequest struct {
	Name          string `json:"name" binding:"required"`
	Slug          string `json:"slug" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Slogan        string `json:"slogan" binding:"required"`
	Domain        string `json:"domain" binding:"required"`
	GitRepoOwner  string `json:"gitRepoOwner" binding:"required"`
	GitRepoName   string `json:"gitRepoName" binding:"required"`
	GitRepoBranch string `json:"gitRepoBranch" binding:"required"`
	GitAPIToken   string `json:"gitApiToken" binding:"required"`
	Config        string `json:"config" binding:"required"`
	LanguageCode  string `json:"languageCode" binding:"required,len=2"`
}

type UpdateWebsiteRequest struct {
	Name          string `json:"name" binding:"omitempty"`
	Slug          string `json:"slug" binding:"omitempty"`
	Description   string `json:"description" binding:"omitempty"`
	Slogan        string `json:"slogan" binding:"omitempty"`
	Domain        string `json:"domain" binding:"omitempty"`
	GitRepoOwner  string `json:"gitRepoOwner" binding:"omitempty"`
	GitRepoName   string `json:"gitRepoName" binding:"omitempty"`
	GitRepoBranch string `json:"gitRepoBranch" binding:"omitempty"`
	GitAPIToken   string `json:"gitApiToken" binding:"omitempty"`
	Config        string `json:"config" binding:"omitempty"`
	LanguageCode  string `json:"languageCode" binding:"omitempty,len=2"`
}
