package study

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/snjwilson/memory-map/internal/core/card"
)

type Service struct {
	reviewRepo ReviewRepository
	cardRepo   card.Repository
}

func NewService(reviewRepo ReviewRepository, cardRepo card.Repository) *Service {
	return &Service{
		reviewRepo: reviewRepo,
		cardRepo:   cardRepo,
	}
}

// ProcessReview handles multiple cards at once
func (s *Service) ProcessReview(ctx context.Context, reviews []Review) ([]*card.Card, error) {
	var updatedCards []*card.Card

	for _, r := range reviews {
		// We reuse your existing logic for each card
		updated, err := s.SubmitReview(ctx, r.CardID, r.Rating, r.Timestamp)
		if err != nil {
			// Log the error but keep processing other cards?
			// Or stop entirely? Usually, we log and continue.
			slog.Error("failed to process card review", "card_id", r.CardID, "error", err)
			continue
		}
		updatedCards = append(updatedCards, updated)
	}

	return updatedCards, nil
}

// SubmitReview processes a user's answer and updates the card
func (s *Service) SubmitReview(ctx context.Context, cardID string, rating Rating, timestamp int64) (*card.Card, error) {
	// 1. Fetch the Card
	card, err := s.cardRepo.GetByID(ctx, cardID)
	if err != nil {
		return nil, err
	}

	// 2. Run the Algorithm (Pure Logic)
	result := CalculateNextReview(card.Interval, card.EaseFactor, card.Repetitions, rating)

	// 3. Update Card State in Memory
	card.Interval = result.Interval
	card.EaseFactor = result.EaseFactor
	card.Repetitions = result.Repetitions

	// Calculate new Due Date: Now + Interval (Days)
	// Note: In a real app, you might want to set this to the *start* of that day (e.g. 00:00)
	card.DueDate = time.Now().UTC().AddDate(0, 0, result.Interval)
	card.UpdatedAt = time.Now().UTC()

	// 4. Save Updates to Database
	if err := s.cardRepo.Update(ctx, card); err != nil {
		return nil, err
	}

	// 5. Create and Save Review Log (Async is optional but good for performance)
	log := &ReviewLog{
		ID:          uuid.New().String(),
		CardID:      card.ID,
		Rating:      rating,
		ReviewTime:  time.UnixMilli(timestamp).UTC(),
		NewInterval: result.Interval,
		NewEase:     result.EaseFactor,
	}

	// We handle the log error gracefully (logging it) rather than failing the whole request
	// In production, use a logger here
	_ = s.reviewRepo.LogReview(ctx, log)

	return card, nil
}
