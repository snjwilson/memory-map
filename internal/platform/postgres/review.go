package postgres

import (
	"context"
	"database/sql"

	"github.com/snjwilson/memory-map/internal/core/study"
)

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

// LogReview inserts a record of a single review event.
func (r *ReviewRepository) LogReview(ctx context.Context, log *study.ReviewLog) error {
	query := `
		INSERT INTO review_logs (
			id, card_id, rating, review_time, duration_ms, new_interval, new_ease
		) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.db.ExecContext(
		ctx,
		query,
		log.ID,
		log.CardID,
		log.Rating,
		log.ReviewTime,
		log.DurationMs,
		log.NewInterval,
		log.NewEase,
	)
	return err
}

// GetRecentReviews can be used to show a "Recent Activity" feed to the user.
func (r *ReviewRepository) GetRecentReviews(ctx context.Context, limit int) ([]*study.ReviewLog, error) {
	query := `
		SELECT id, card_id, rating, review_time, duration_ms, new_interval, new_ease
		FROM review_logs
		ORDER BY review_time DESC
		LIMIT $1`

	rows, err := r.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*study.ReviewLog
	for rows.Next() {
		l := &study.ReviewLog{}
		err := rows.Scan(
			&l.ID,
			&l.CardID,
			&l.Rating,
			&l.ReviewTime,
			&l.DurationMs,
			&l.NewInterval,
			&l.NewEase,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, l)
	}
	return logs, nil
}
