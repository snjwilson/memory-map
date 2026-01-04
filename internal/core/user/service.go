package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// SignUp handles new user registration and returns an auth token
func (s *Service) SignUp(ctx context.Context, req SignUpRequest) (*AuthToken, error) {
	// 1. Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 2. Prepare the entity
	now := time.Now().UTC()
	u := &User{
		ID:           uuid.New().String(),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	// 3. Persist
	if err := s.repo.Create(ctx, u); err != nil {
		// In production, check if error is a unique constraint violation
		return nil, err
	}

	// 4. Generate JWT token
	token, err := GenerateToken(u)
	if err != nil {
		return nil, err
	}

	return &AuthToken{User: u, Token: token}, nil
}

// Login verifies credentials and returns an auth token
func (s *Service) Login(ctx context.Context, email, password string) (*AuthToken, error) {
	u, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Compare hashed password with plain text
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := GenerateToken(u)
	if err != nil {
		return nil, err
	}

	return &AuthToken{User: u, Token: token}, nil
}
