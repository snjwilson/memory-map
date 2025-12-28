package deck

import (
	"context"
)

// Repository defines how decks are stored and retrieved.
// The actual implementation (Postgres/SQL) will live in /internal/platform/postgres
type Repository interface {
	// Create saves a new deck and returns the ID
	Create(ctx context.Context, deck *Deck) error

	// GetByID retrieves a specific deck
	GetByID(ctx context.Context, id string) (*Deck, error)

	// ListByOwner retrieves all decks for a specific user
	ListByOwner(ctx context.Context, ownerID string) ([]*Deck, error)

	// Update modifies an existing deck
	Update(ctx context.Context, deck *Deck) error

	// Delete removes a deck
	Delete(ctx context.Context, id string) error
}
