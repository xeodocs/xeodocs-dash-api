package models

import (
	"time"
)

type Page struct {
	ID                   int        `json:"id" db:"id"`
	WebsiteID            int        `json:"websiteId" db:"website_id"`
	Title                string     `json:"title" db:"title"`
	Slug                 string     `json:"slug" db:"slug"`
	Description          string     `json:"description" db:"description"`
	MarkdownContent      string     `json:"markdownContent" db:"markdown_content"`
	Tags                 string     `json:"tags" db:"tags"`
	FreezeStatus         bool       `json:"freezeStatus" db:"freeze_status"`
	Status               string     `json:"status" db:"status"`
	LastStatusChangeAt   time.Time  `json:"lastStatusChangeAt" db:"last_status_change_at"`
	ScheduledPublishAt   *time.Time `json:"scheduledPublishAt" db:"scheduled_publish_at"`
	CreatedAt            time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt            time.Time  `json:"updatedAt" db:"updated_at"`
}

type PageAsset struct {
	ID        int       `json:"id" db:"id"`
	PageID    int       `json:"pageId" db:"page_id"`
	BucketKey string    `json:"bucketKey" db:"bucket_key"`
	MimeType  string    `json:"mimeType" db:"mime_type"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

// Request/Response DTOs
type CreatePageRequest struct {
	WebsiteID           int        `json:"websiteId" binding:"required"`
	Title               string     `json:"title" binding:"required"`
	Slug                string     `json:"slug" binding:"required"`
	Description         string     `json:"description" binding:"required"`
	MarkdownContent     string     `json:"markdownContent" binding:"required"`
	Tags                string     `json:"tags" binding:"omitempty"`
	FreezeStatus        bool       `json:"freezeStatus"`
	Status              string     `json:"status" binding:"required,oneof=draft translating translated ignored published"`
	ScheduledPublishAt  *time.Time `json:"scheduledPublishAt"`
}

type UpdatePageRequest struct {
	Title               string     `json:"title" binding:"omitempty"`
	Slug                string     `json:"slug" binding:"omitempty"`
	Description         string     `json:"description" binding:"omitempty"`
	MarkdownContent     string     `json:"markdownContent" binding:"omitempty"`
	Tags                string     `json:"tags" binding:"omitempty"`
	FreezeStatus        *bool      `json:"freezeStatus"`
	Status              string     `json:"status" binding:"omitempty,oneof=draft translating translated ignored published"`
	ScheduledPublishAt  *time.Time `json:"scheduledPublishAt"`
}
