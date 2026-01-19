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
	OwnerID     string    `json:"ownerId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"isPublic"`
	CardCount   int       `db:"card_count" json:"cardCount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// NewDeckRequest defines the data needed to create a deck
type NewDeckRequest struct {
	OwnerID     string `json:"ownerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPublic    bool   `json:"isPublic"`
}
