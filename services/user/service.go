package main

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Service errors
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists        = errors.New("user already exists")
)

// User represents the user model
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

// Service defines the user service interface
type Service interface {
	CreateUser(ctx context.Context, username, password, email string) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	ValidateUser(ctx context.Context, username, password string) (*User, error)
}

type service struct {
	users map[string]*User
}

// NewService creates a new user service instance
func NewService() Service {
	return &service{
		users: make(map[string]*User),
	}
}

// CreateUser creates a new user
func (s *service) CreateUser(ctx context.Context, username, password, email string) (*User, error) {
	// Check if user already exists
	for _, u := range s.users {
		if u.Username == username {
			return nil, ErrUserExists
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       uuid.New().String(),
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	s.users[user.ID] = user
	return user, nil
}

// GetUser retrieves a user by ID
func (s *service) GetUser(ctx context.Context, id string) (*User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

// ValidateUser validates user credentials
func (s *service) ValidateUser(ctx context.Context, username, password string) (*User, error) {
	for _, user := range s.users {
		if user.Username == username {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err == nil {
				return user, nil
			}
			return nil, ErrInvalidCredentials
		}
	}
	return nil, ErrUserNotFound
}
