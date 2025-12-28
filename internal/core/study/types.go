package study

import (
	"time"
)

// Rating represents how well the user remembered the card
type Rating int

const (
	RatingAgain Rating = 1 // Complete blackout, need to see immediately
	RatingHard  Rating = 2 // Remembered, but with great difficulty
	RatingGood  Rating = 3 // Correct response with some hesitation
	RatingEasy  Rating = 4 // Perfect recall
)

// ReviewLog records a single practice event.
// This is crucial for analytics (e.g., "Heatmap" or "Cards Learned per Day")
type ReviewLog struct {
	ID         string    `json:"id"`
	CardID     string    `json:"card_id"`
	Rating     Rating    `json:"rating"`
	ReviewTime time.Time `json:"review_time"` // When it happened
	DurationMs int       `json:"duration_ms"` // How long they took to answer

	// Snapshot of state *after* the review (optional, but helpful for debugging)
	NewInterval int     `json:"new_interval"`
	NewEase     float64 `json:"new_ease"`
}
