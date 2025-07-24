package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/xeodocs/xeodocs-dash-api/internal/models"
)

type WebsiteRepository struct {
	db *sql.DB
}

func NewWebsiteRepository(db *sql.DB) *WebsiteRepository {
	return &WebsiteRepository{db: db}
}

func (r *WebsiteRepository) Create(website *models.Website) error {
	query := `
		INSERT INTO websites (name, slug, description, slogan, domain, git_repo_owner, 
			git_repo_name, git_repo_branch, git_api_token, config, language_code, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	now := time.Now()
	result, err := r.db.Exec(query, website.Name, website.Slug, website.Description,
		website.Slogan, website.Domain, website.GitRepoOwner, website.GitRepoName,
		website.GitRepoBranch, website.GitAPIToken, website.Config, website.LanguageCode,
		now, now)
	if err != nil {
		return fmt.Errorf("failed to create website: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get website ID: %w", err)
	}

	website.ID = int(id)
	website.CreatedAt = now
	website.UpdatedAt = now
	return nil
}

func (r *WebsiteRepository) GetByID(id int) (*models.Website, error) {
	query := `
		SELECT id, name, slug, description, slogan, domain, git_repo_owner, 
			git_repo_name, git_repo_branch, git_api_token, config, language_code, created_at, updated_at
		FROM websites WHERE id = ?
	`
	website := &models.Website{}
	err := r.db.QueryRow(query, id).Scan(
		&website.ID, &website.Name, &website.Slug, &website.Description,
		&website.Slogan, &website.Domain, &website.GitRepoOwner, &website.GitRepoName,
		&website.GitRepoBranch, &website.GitAPIToken, &website.Config, &website.LanguageCode,
		&website.CreatedAt, &website.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("website not found")
		}
		return nil, fmt.Errorf("failed to get website: %w", err)
	}
	return website, nil
}

func (r *WebsiteRepository) GetBySlug(slug string) (*models.Website, error) {
	query := `
		SELECT id, name, slug, description, slogan, domain, git_repo_owner, 
			git_repo_name, git_repo_branch, git_api_token, config, language_code, created_at, updated_at
		FROM websites WHERE slug = ?
	`
	website := &models.Website{}
	err := r.db.QueryRow(query, slug).Scan(
		&website.ID, &website.Name, &website.Slug, &website.Description,
		&website.Slogan, &website.Domain, &website.GitRepoOwner, &website.GitRepoName,
		&website.GitRepoBranch, &website.GitAPIToken, &website.Config, &website.LanguageCode,
		&website.CreatedAt, &website.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("website not found")
		}
		return nil, fmt.Errorf("failed to get website: %w", err)
	}
	return website, nil
}

func (r *WebsiteRepository) GetAll() ([]*models.Website, error) {
	query := `
		SELECT id, name, slug, description, slogan, domain, git_repo_owner, 
			git_repo_name, git_repo_branch, git_api_token, config, language_code, created_at, updated_at
		FROM websites ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get websites: %w", err)
	}
	defer rows.Close()

	var websites []*models.Website
	for rows.Next() {
		website := &models.Website{}
		err := rows.Scan(
			&website.ID, &website.Name, &website.Slug, &website.Description,
			&website.Slogan, &website.Domain, &website.GitRepoOwner, &website.GitRepoName,
			&website.GitRepoBranch, &website.GitAPIToken, &website.Config, &website.LanguageCode,
			&website.CreatedAt, &website.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan website: %w", err)
		}
		websites = append(websites, website)
	}
	return websites, nil
}

func (r *WebsiteRepository) Update(id int, website *models.Website) error {
	query := `
		UPDATE websites SET name = ?, slug = ?, description = ?, slogan = ?, domain = ?, 
			git_repo_owner = ?, git_repo_name = ?, git_repo_branch = ?, git_api_token = ?, 
			config = ?, language_code = ?, updated_at = ?
		WHERE id = ?
	`
	now := time.Now()
	result, err := r.db.Exec(query, website.Name, website.Slug, website.Description,
		website.Slogan, website.Domain, website.GitRepoOwner, website.GitRepoName,
		website.GitRepoBranch, website.GitAPIToken, website.Config, website.LanguageCode,
		now, id)
	if err != nil {
		return fmt.Errorf("failed to update website: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("website not found")
	}

	website.UpdatedAt = now
	return nil
}

func (r *WebsiteRepository) Delete(id int) error {
	query := `DELETE FROM websites WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete website: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("website not found")
	}

	return nil
}
