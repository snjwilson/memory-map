package card

import (
	"context"
)

type Repository interface {
	// Create saves a new card
	Create(ctx context.Context, card *Card) error

	// GetByID retrieves a specific card
	GetByID(ctx context.Context, id string) (*Card, error)

	// GetByDeckID retrieves all cards in a deck (for management view)
	GetByDeckID(ctx context.Context, deckID string) ([]*Card, error)

	// GetDueCards retrieves cards where DueDate <= Now (for study session)
	GetDueCards(ctx context.Context, deckID string, limit int) ([]*Card, error)

	// Update saves changes (content edits OR algorithm updates)
	Update(ctx context.Context, card *Card) error

	// Delete removes a card
	Delete(ctx context.Context, id string) error
}
