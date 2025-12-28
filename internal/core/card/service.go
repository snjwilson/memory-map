package card

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

// DeckValidator ensures we don't import the full 'decks' package
// We just need a way to check if a deck exists/belongs to the user.
type DeckValidator interface {
	Exists(ctx context.Context, deckID string) (bool, error)
}

type Service struct {
	repo          Repository
	deckValidator DeckValidator
}

// NewService creates the card service
func NewService(repo Repository, validator DeckValidator) *Service {
	return &Service{
		repo:          repo,
		deckValidator: validator,
	}
}

func (s *Service) CreateCard(ctx context.Context, req NewCardRequest) (*Card, error) {
	// 1. Validate Content
	if req.Front == "" {
		return nil, ErrInvalidContent
	}

	// 2. Validate Deck Exists (Using the interface)
	exists, err := s.deckValidator.Exists(ctx, req.DeckID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("deck does not exist")
	}

	// 3. Initialize default learning state
	now := time.Now().UTC()
	card := &Card{
		ID:     uuid.New().String(),
		DeckID: req.DeckID,
		Front:  req.Front,
		Back:   req.Back,

		// SM-2 Defaults
		Interval:    0,   // 0 means it's "New"
		EaseFactor:  2.5, // Standard starting ease
		Repetitions: 0,
		DueDate:     now, // Due immediately

		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(ctx, card); err != nil {
		return nil, err
	}

	return card, nil
}

// GetDueCards is used when the user hits "Study"
func (s *Service) GetDueCards(ctx context.Context, deckID string) ([]*Card, error) {
	// We might limit the batch size here (e.g., fetch 20 at a time)
	return s.repo.GetDueCards(ctx, deckID, 20)
}
