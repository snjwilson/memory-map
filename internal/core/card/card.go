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
	DeckID string `json:"deckId"`
	Front  string `json:"front"` // The Question
	Back   string `json:"back"`  // The Answer

	// Learning State (SM-2 Algorithm Data)
	// These are modified by the study package, but stored here.
	Interval    int       `json:"interval"`    // Days until next review
	EaseFactor  float64   `json:"easeFactor"`  // Multiplier (default 2.5)
	Repetitions int       `json:"repetitions"` // Consecutive successful recalls
	DueDate     time.Time `json:"dueDate"`     // When to show this card next

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewCardRequest is the input for creating a card
type NewCardRequest struct {
	DeckID string `json:"deckId"`
	Front  string `json:"front"`
	Back   string `json:"back"`
}
