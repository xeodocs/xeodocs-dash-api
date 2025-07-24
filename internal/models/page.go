package models

import (
	"time"
)

type Page struct {
	ID                   int        `json:"id" db:"id"`
	WebsiteID            int        `json:"website_id" db:"website_id"`
	Title                string     `json:"title" db:"title"`
	Slug                 string     `json:"slug" db:"slug"`
	Description          string     `json:"description" db:"description"`
	MarkdownContent      string     `json:"markdown_content" db:"markdown_content"`
	Tags                 string     `json:"tags" db:"tags"`
	FreezeStatus         bool       `json:"freeze_status" db:"freeze_status"`
	Status               string     `json:"status" db:"status"`
	LastStatusChangeAt   time.Time  `json:"last_status_change_at" db:"last_status_change_at"`
	ScheduledPublishAt   *time.Time `json:"scheduled_publish_at" db:"scheduled_publish_at"`
	CreatedAt            time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at" db:"updated_at"`
}

type PageAsset struct {
	ID        int       `json:"id" db:"id"`
	PageID    int       `json:"page_id" db:"page_id"`
	BucketKey string    `json:"bucket_key" db:"bucket_key"`
	MimeType  string    `json:"mime_type" db:"mime_type"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Request/Response DTOs
type CreatePageRequest struct {
	WebsiteID           int        `json:"website_id" binding:"required"`
	Title               string     `json:"title" binding:"required"`
	Slug                string     `json:"slug" binding:"required"`
	Description         string     `json:"description" binding:"required"`
	MarkdownContent     string     `json:"markdown_content" binding:"required"`
	Tags                string     `json:"tags" binding:"omitempty"`
	FreezeStatus        bool       `json:"freeze_status"`
	Status              string     `json:"status" binding:"required,oneof=draft translating translated ignored published"`
	ScheduledPublishAt  *time.Time `json:"scheduled_publish_at"`
}

type UpdatePageRequest struct {
	Title               string     `json:"title" binding:"omitempty"`
	Slug                string     `json:"slug" binding:"omitempty"`
	Description         string     `json:"description" binding:"omitempty"`
	MarkdownContent     string     `json:"markdown_content" binding:"omitempty"`
	Tags                string     `json:"tags" binding:"omitempty"`
	FreezeStatus        *bool      `json:"freeze_status"`
	Status              string     `json:"status" binding:"omitempty,oneof=draft translating translated ignored published"`
	ScheduledPublishAt  *time.Time `json:"scheduled_publish_at"`
}
