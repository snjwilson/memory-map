package study

import (
	"time"
)

// Rating represents how well the user remembered the card
type Rating int

const (
	RatingAgain Rating = 1 // Wrong
	RatingHard  Rating = 2 // Correct, but tough
	RatingEasy  Rating = 3 // Correct, easy
)

// ReviewLog records a single practice event.
// This is crucial for analytics (e.g., "Heatmap" or "Cards Learned per Day")
type ReviewLog struct {
	ID         string    `json:"id"`
	CardID     string    `json:"cardId"`
	Rating     Rating    `json:"rating"`
	ReviewTime time.Time `json:"reviewTime"` // When it happened
	DurationMs int       `json:"durationMs"` // How long they took to answer

	// Snapshot of state *after* the review (optional, but helpful for debugging)
	NewInterval int     `json:"newInterval"`
	NewEase     float64 `json:"newEase"`
}

type Review struct {
	CardID    string `json:"cardId"`
	Rating    Rating `json:"grade"` // 1-3
	Timestamp int64  `json:"timestamp"`
}

type StudySessionReviewRequest struct {
	Reviews []Review `json:"reviews"` // Changed to a slice
}
