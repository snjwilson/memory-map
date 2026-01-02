package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/snjwilson/memory-map/internal/core/user"
)

// UserRepository handles persistence logic for Users using PostgreSQL
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository returns a new instance of the postgres repository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Save inserts a new user record into the database
func (r *UserRepository) Create(ctx context.Context, user *user.User) error {
	query := `
		INSERT INTO users (id, first_name, last_name, email, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.FirstName, user.LastName, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error saving user: %w", err)
	}
	return nil
}

// GetByID retrieves a user by their unique identifier
func (r *UserRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	query := `SELECT id, first_name, last_name, email, password_hash, created_at, updated_at FROM users WHERE id = $1`

	u := &user.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PasswordHash,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u, user.ErrUserNotFound
		}
		return u, fmt.Errorf("error fetching user: %w", err)
	}

	return u, nil
}

// GetByID retrieves a user by their unique identifier
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `SELECT id, first_name, last_name, email, password_hash, created_at, updated_at FROM users WHERE email = $1`

	u := &user.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PasswordHash,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u, user.ErrUserNotFound
		}
		return u, fmt.Errorf("error fetching user: %w", err)
	}

	return u, nil
}
