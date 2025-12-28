package card

import (
	"errors"
	"time"
)

var (
	ErrCardNotFound   = errors.New("card not found")
	ErrInvalidContent = errors.New("card front cannot be empty")
)

// Card represents a single flashcard
type Card struct {
	ID     string `json:"id"`
	DeckID string `json:"deck_id"`
	Front  string `json:"front"` // The Question
	Back   string `json:"back"`  // The Answer

	// Learning State (SM-2 Algorithm Data)
	// These are modified by the study package, but stored here.
	Interval    int       `json:"interval"`    // Days until next review
	EaseFactor  float64   `json:"ease_factor"` // Multiplier (default 2.5)
	Repetitions int       `json:"repetitions"` // Consecutive successful recalls
	DueDate     time.Time `json:"due_date"`    // When to show this card next

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewCardRequest is the input for creating a card
type NewCardRequest struct {
	DeckID string
	Front  string
	Back   string
}
