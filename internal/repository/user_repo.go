package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/xeodocs/xeodocs-dash-api/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (email, password_hash, name, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`
	now := time.Now()
	result, err := r.db.Exec(query, user.Email, user.PasswordHash, user.Name, now, now)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get user ID: %w", err)
	}

	user.ID = int(id)
	user.CreatedAt = now
	user.UpdatedAt = now
	return nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	query := `
		SELECT id, email, password_hash, name, created_at, updated_at
		FROM users WHERE id = ?
	`
	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Name,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, email, password_hash, name, created_at, updated_at
		FROM users WHERE email = ?
	`
	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Name,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (r *UserRepository) GetAll() ([]*models.User, error) {
	query := `
		SELECT id, email, password_hash, name, created_at, updated_at
		FROM users ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID, &user.Email, &user.PasswordHash, &user.Name,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Update(id int, user *models.User) error {
	query := `
		UPDATE users SET email = ?, name = ?, updated_at = ?
		WHERE id = ?
	`
	now := time.Now()
	result, err := r.db.Exec(query, user.Email, user.Name, now, id)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	user.UpdatedAt = now
	return nil
}

func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// Session management methods
func (r *UserRepository) CreateSession(session *models.UserSession) error {
	query := `
		INSERT INTO user_sessions (user_id, session_token, expires_at, created_at)
		VALUES (?, ?, ?, ?)
	`
	now := time.Now()
	result, err := r.db.Exec(query, session.UserID, session.SessionToken, session.ExpiresAt, now)
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get session ID: %w", err)
	}

	session.ID = int(id)
	session.CreatedAt = now
	return nil
}

func (r *UserRepository) GetSessionByToken(token string) (*models.UserSession, error) {
	query := `
		SELECT id, user_id, session_token, expires_at, created_at
		FROM user_sessions WHERE session_token = ? AND expires_at > ?
	`
	session := &models.UserSession{}
	err := r.db.QueryRow(query, token, time.Now()).Scan(
		&session.ID, &session.UserID, &session.SessionToken,
		&session.ExpiresAt, &session.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("session not found or expired")
		}
		return nil, fmt.Errorf("failed to get session: %w", err)
	}
	return session, nil
}

func (r *UserRepository) DeleteSession(token string) error {
	query := `DELETE FROM user_sessions WHERE session_token = ?`
	_, err := r.db.Exec(query, token)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}
	return nil
}

func (r *UserRepository) DeleteExpiredSessions() error {
	query := `DELETE FROM user_sessions WHERE expires_at <= ?`
	_, err := r.db.Exec(query, time.Now())
	if err != nil {
		return fmt.Errorf("failed to delete expired sessions: %w", err)
	}
	return nil
}
