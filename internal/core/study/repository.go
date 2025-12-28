package study

import "context"

type ReviewRepository interface {
	// LogReview saves the history of a review
	LogReview(ctx context.Context, log *ReviewLog) error

	// GetRecentReviews could be used for an "Undo" feature or analytics
	GetRecentReviews(ctx context.Context, limit int) ([]*ReviewLog, error)
}
