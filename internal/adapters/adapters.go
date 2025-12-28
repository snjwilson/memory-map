package adapters

import (
	"context"

	"github.com/snjwilson/memory-map/internal/core/deck"
)

// This struct makes the Decks Service look like a Deck Validator
type DeckServiceAdapter struct {
	Service *deck.Service
}

func (a *DeckServiceAdapter) Exists(ctx context.Context, id string) (bool, error) {
	deck, err := a.Service.GetDeckById(ctx, id)
	if err != nil {
		// Handle "not found" vs "db error"
		return false, err
	}
	return deck != nil, nil
}
