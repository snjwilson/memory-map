package deck

import (
	"context"
	"time"

	"github.com/google/uuid" // You might need: go get github.com/google/uuid
)

// Service handles all business logic for decks
type Service struct {
	repo Repository
}

// NewService creates a deck service with the necessary dependencies
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// CreateDeck validates input and persists a new deck
func (s *Service) CreateDeck(ctx context.Context, req NewDeckRequest) (*Deck, error) {
	// 1. Validation Logic
	if req.Name == "" {
		return nil, ErrInvalidDeck
	}

	// 2. Prepare the entity
	now := time.Now().UTC()
	deck := &Deck{
		ID:          uuid.New().String(), // Generate a UUID
		OwnerID:     req.OwnerID,
		Name:        req.Name,
		Description: req.Description,
		IsPublic:    req.IsPublic,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// 3. Persist using the interface
	if err := s.repo.Create(ctx, deck); err != nil {
		return nil, err
	}

	return deck, nil
}

// GetUserDecks fetches all decks for a user
func (s *Service) GetUserDecks(ctx context.Context, userID string) ([]*Deck, error) {
	return s.repo.ListByOwner(ctx, userID)
}

// GetDeckById fetches one deck by id
// returns an error if deck does not exist
func (s *Service) GetDeckById(ctx context.Context, id string) (*Deck, error) {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

// UpdateDeck handles renaming or changing description
func (s *Service) UpdateDeck(ctx context.Context, id string, name string, description string) error {
	// 1. Fetch existing to ensure it exists
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// 2. Apply updates
	existing.Name = name
	existing.Description = description
	existing.UpdatedAt = time.Now().UTC()

	// 3. Save
	return s.repo.Update(ctx, existing)
}
