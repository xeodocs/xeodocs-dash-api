package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/xeodocs/xeodocs-dash-api/internal/models"
	"github.com/xeodocs/xeodocs-dash-api/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Name:         req.Name,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	return s.userRepo.GetAll()
}

func (s *UserService) UpdateUser(id int, req *models.UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Email != "" {
		// Check if new email is already taken by another user
		existingUser, _ := s.userRepo.GetByEmail(req.Email)
		if existingUser != nil && existingUser.ID != id {
			return nil, fmt.Errorf("email %s is already taken", req.Email)
		}
		user.Email = req.Email
	}
	if req.Name != "" {
		user.Name = req.Name
	}

	err = s.userRepo.Update(id, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) Login(req *models.LoginRequest) (*models.LoginResponse, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Generate session token
	sessionToken, err := s.generateSessionToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate session token: %w", err)
	}

	// Create session
	session := &models.UserSession{
		UserID:       user.ID,
		SessionToken: sessionToken,
		ExpiresAt:    time.Now().Add(24 * time.Hour), // 24 hours expiry
	}

	err = s.userRepo.CreateSession(session)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return &models.LoginResponse{
		User:         *user,
		SessionToken: sessionToken,
	}, nil
}

func (s *UserService) Logout(sessionToken string) error {
	return s.userRepo.DeleteSession(sessionToken)
}

func (s *UserService) ValidateSession(sessionToken string) (*models.User, error) {
	session, err := s.userRepo.GetSessionByToken(sessionToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetByID(session.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) CleanupExpiredSessions() error {
	return s.userRepo.DeleteExpiredSessions()
}

func (s *UserService) generateSessionToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
