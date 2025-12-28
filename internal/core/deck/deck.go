package deck

import (
	"errors"
	"time"
)

// Common errors for this domain
var (
	ErrDeckNotFound = errors.New("deck not found")
	ErrInvalidDeck  = errors.New("deck name is required")
)

// Deck represents a collection of cards
type Deck struct {
	ID          string    `json:"id"`
	OwnerID     string    `json:"owner_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"is_public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewDeckRequest defines the data needed to create a deck
type NewDeckRequest struct {
	OwnerID     string
	Name        string
	Description string
	IsPublic    bool
}
