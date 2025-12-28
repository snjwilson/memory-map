package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/snjwilson/memory-map/internal/core/card"
)

type CardRepository struct {
	db *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{db: db}
}

// GetDueCards finds cards that are ready for review
func (r *CardRepository) GetDueCards(ctx context.Context, deckID string, limit int) ([]*card.Card, error) {
	query := `SELECT id, deck_id, front, back, interval, ease_factor, repetitions, due_date 
              FROM cards 
              WHERE deck_id = $1 AND due_date <= $2 
              ORDER BY due_date ASC 
              LIMIT $3`

	rows, err := r.db.QueryContext(ctx, query, deckID, time.Now().UTC(), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*card.Card
	for rows.Next() {
		c := &card.Card{}
		err := rows.Scan(&c.ID, &c.DeckID, &c.Front, &c.Back, &c.Interval, &c.EaseFactor, &c.Repetitions, &c.DueDate)
		if err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}

// Create saves a new card
func (r *CardRepository) Create(ctx context.Context, c *card.Card) error {
	query := `INSERT INTO cards (id, deck_id, front, back, interval, ease_factor, repetitions, due_date, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.db.ExecContext(ctx, query, c.ID, c.DeckID, c.Front, c.Back, c.Interval, c.EaseFactor, c.Repetitions, c.DueDate, time.Now().UTC(), time.Now().UTC())
	return err
}

// GetByID retrieves a specific card
func (r *CardRepository) GetByID(ctx context.Context, id string) (*card.Card, error) {
	c := &card.Card{}
	query := `SELECT id, deck_id, front, back, interval, ease_factor, repetitions, due_date FROM cards WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.DeckID, &c.Front, &c.Back, &c.Interval, &c.EaseFactor, &c.Repetitions, &c.DueDate)
	if err == sql.ErrNoRows {
		return nil, card.ErrCardNotFound
	}
	return c, err
}

// GetByDeckID retrieves all cards in a deck (for management view)
func (r *CardRepository) GetByDeckID(ctx context.Context, deckID string) ([]*card.Card, error) {
	query := `SELECT id, deck_id, front, back, interval, ease_factor, repetitions, due_date FROM cards WHERE deck_id = $1`
	rows, err := r.db.QueryContext(ctx, query, deckID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*card.Card
	for rows.Next() {
		c := &card.Card{}
		err := rows.Scan(&c.ID, &c.DeckID, &c.Front, &c.Back, &c.Interval, &c.EaseFactor, &c.Repetitions, &c.DueDate)
		if err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}

// Update is used to save the new Interval/EaseFactor after a study session
func (r *CardRepository) Update(ctx context.Context, c *card.Card) error {
	query := `UPDATE cards SET 
                front = $1, back = $2, interval = $3, 
                ease_factor = $4, repetitions = $5, due_date = $6, updated_at = $7 
              WHERE id = $8`
	_, err := r.db.ExecContext(ctx, query,
		c.Front, c.Back, c.Interval,
		c.EaseFactor, c.Repetitions, c.DueDate, time.Now().UTC(),
		c.ID,
	)
	return err
}

// Delete removes a card
func (r *CardRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM cards WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
