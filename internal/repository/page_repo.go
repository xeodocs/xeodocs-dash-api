package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/xeodocs/xeodocs-dash-api/internal/models"
)

type PageRepository struct {
	db *sql.DB
}

func NewPageRepository(db *sql.DB) *PageRepository {
	return &PageRepository{db: db}
}

func (r *PageRepository) Create(page *models.Page) error {
	query := `
		INSERT INTO pages (website_id, title, slug, description, markdown_content, tags, 
			freeze_status, status, last_status_change_at, scheduled_publish_at, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	now := time.Now()
	result, err := r.db.Exec(query, page.WebsiteID, page.Title, page.Slug, page.Description,
		page.MarkdownContent, page.Tags, page.FreezeStatus, page.Status, now,
		page.ScheduledPublishAt, now, now)
	if err != nil {
		return fmt.Errorf("failed to create page: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get page ID: %w", err)
	}

	page.ID = int(id)
	page.LastStatusChangeAt = now
	page.CreatedAt = now
	page.UpdatedAt = now
	return nil
}

func (r *PageRepository) GetByID(id int) (*models.Page, error) {
	query := `
		SELECT id, website_id, title, slug, description, markdown_content, tags, 
			freeze_status, status, last_status_change_at, scheduled_publish_at, created_at, updated_at
		FROM pages WHERE id = ?
	`
	page := &models.Page{}
	err := r.db.QueryRow(query, id).Scan(
		&page.ID, &page.WebsiteID, &page.Title, &page.Slug, &page.Description,
		&page.MarkdownContent, &page.Tags, &page.FreezeStatus, &page.Status,
		&page.LastStatusChangeAt, &page.ScheduledPublishAt, &page.CreatedAt, &page.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("page not found")
		}
		return nil, fmt.Errorf("failed to get page: %w", err)
	}
	return page, nil
}

func (r *PageRepository) GetBySlug(slug string) (*models.Page, error) {
	query := `
		SELECT id, website_id, title, slug, description, markdown_content, tags, 
			freeze_status, status, last_status_change_at, scheduled_publish_at, created_at, updated_at
		FROM pages WHERE slug = ?
	`
	page := &models.Page{}
	err := r.db.QueryRow(query, slug).Scan(
		&page.ID, &page.WebsiteID, &page.Title, &page.Slug, &page.Description,
		&page.MarkdownContent, &page.Tags, &page.FreezeStatus, &page.Status,
		&page.LastStatusChangeAt, &page.ScheduledPublishAt, &page.CreatedAt, &page.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("page not found")
		}
		return nil, fmt.Errorf("failed to get page: %w", err)
	}
	return page, nil
}

func (r *PageRepository) GetAll() ([]*models.Page, error) {
	query := `
		SELECT id, website_id, title, slug, description, markdown_content, tags, 
			freeze_status, status, last_status_change_at, scheduled_publish_at, created_at, updated_at
		FROM pages ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get pages: %w", err)
	}
	defer rows.Close()

	var pages []*models.Page
	for rows.Next() {
		page := &models.Page{}
		err := rows.Scan(
			&page.ID, &page.WebsiteID, &page.Title, &page.Slug, &page.Description,
			&page.MarkdownContent, &page.Tags, &page.FreezeStatus, &page.Status,
			&page.LastStatusChangeAt, &page.ScheduledPublishAt, &page.CreatedAt, &page.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan page: %w", err)
		}
		pages = append(pages, page)
	}
	return pages, nil
}

func (r *PageRepository) GetByWebsiteID(websiteID int) ([]*models.Page, error) {
	query := `
		SELECT id, website_id, title, slug, description, markdown_content, tags, 
			freeze_status, status, last_status_change_at, scheduled_publish_at, created_at, updated_at
		FROM pages WHERE website_id = ? ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query, websiteID)
	if err != nil {
		return nil, fmt.Errorf("failed to get pages by website: %w", err)
	}
	defer rows.Close()

	var pages []*models.Page
	for rows.Next() {
		page := &models.Page{}
		err := rows.Scan(
			&page.ID, &page.WebsiteID, &page.Title, &page.Slug, &page.Description,
			&page.MarkdownContent, &page.Tags, &page.FreezeStatus, &page.Status,
			&page.LastStatusChangeAt, &page.ScheduledPublishAt, &page.CreatedAt, &page.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan page: %w", err)
		}
		pages = append(pages, page)
	}
	return pages, nil
}

func (r *PageRepository) Update(id int, page *models.Page) error {
	query := `
		UPDATE pages SET title = ?, slug = ?, description = ?, markdown_content = ?, 
			tags = ?, freeze_status = ?, status = ?, last_status_change_at = ?, 
			scheduled_publish_at = ?, updated_at = ?
		WHERE id = ?
	`
	now := time.Now()
	result, err := r.db.Exec(query, page.Title, page.Slug, page.Description,
		page.MarkdownContent, page.Tags, page.FreezeStatus, page.Status,
		page.LastStatusChangeAt, page.ScheduledPublishAt, now, id)
	if err != nil {
		return fmt.Errorf("failed to update page: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("page not found")
	}

	page.UpdatedAt = now
	return nil
}

func (r *PageRepository) Delete(id int) error {
	query := `DELETE FROM pages WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete page: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("page not found")
	}

	return nil
}
