package models

import (
	"time"
)

type User struct {
	ID           int       `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Name         string    `json:"name" db:"name"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type UserSession struct {
	ID           int       `json:"id" db:"id"`
	UserID       int       `json:"user_id" db:"user_id"`
	SessionToken string    `json:"session_token" db:"session_token"`
	ExpiresAt    time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type Role struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Permissions string    `json:"permissions" db:"permissions"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type UserRole struct {
	UserID int `json:"user_id" db:"user_id"`
	RoleID int `json:"role_id" db:"role_id"`
}

// Request/Response DTOs
type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

type UpdateUserRequest struct {
	Email string `json:"email" binding:"omitempty,email"`
	Name  string `json:"name" binding:"omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	User         User   `json:"user"`
	SessionToken string `json:"session_token"`
}
