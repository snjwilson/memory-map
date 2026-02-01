package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/snjwilson/memory-map/internal/core/deck"
)

type DeckRepository struct {
	db *sql.DB
}

func NewDeckRepository(db *sql.DB) *DeckRepository {
	return &DeckRepository{db: db}
}

func (r *DeckRepository) Create(ctx context.Context, d *deck.Deck) error {
	query := `INSERT INTO decks (id, owner_id, name, description, is_public, card_count, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.ExecContext(ctx, query, d.ID, d.OwnerID, d.Name, d.Description, d.IsPublic, d.CardCount, d.CreatedAt, d.UpdatedAt)
	return err
}

func (r *DeckRepository) GetByID(ctx context.Context, id string) (*deck.Deck, error) {
	d := &deck.Deck{}
	query := `SELECT id, owner_id, name, description, is_public, card_count, created_at, updated_at FROM decks WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&d.ID, &d.OwnerID, &d.Name, &d.Description, &d.IsPublic, &d.CardCount, &d.CreatedAt, &d.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, deck.ErrDeckNotFound
	}
	return d, err
}

// ListByOwner returns the slice of decks, the total count, and any error
func (r *DeckRepository) ListByOwner(ctx context.Context, ownerId string, page, limit int) ([]*deck.Deck, int, error) {
	offset := (page - 1) * limit

	// 1. Get the Total Count
	var totalCount int
	countQuery := `SELECT COUNT(*) FROM decks WHERE owner_id = $1`
	err := r.db.QueryRowContext(ctx, countQuery, ownerId).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count decks: %w", err)
	}

	// If no records exist, we can skip the second query
	if totalCount == 0 {
		return []*deck.Deck{}, 0, nil
	}

	// 2. Get the Paginated Data
	query := `SELECT id, owner_id, name, description, is_public, card_count, created_at, updated_at 
              FROM decks 
              WHERE owner_id = $1 
              ORDER BY created_at DESC 
              LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, ownerId, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query decks: %w", err)
	}
	defer rows.Close()

	result := []*deck.Deck{}
	for rows.Next() {
		d := &deck.Deck{}
		err := rows.Scan(&d.ID, &d.OwnerID, &d.Name, &d.Description, &d.IsPublic, &d.CardCount, &d.CreatedAt, &d.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}
		result = append(result, d)
	}

	return result, totalCount, nil
}

func (r *DeckRepository) Update(ctx context.Context, d *deck.Deck) error {
	query := `UPDATE decks SET name = $1, description = $2, is_public = $3, updated_at = $4 WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query, d.Name, d.Description, d.IsPublic, d.UpdatedAt, d.ID)
	return err
}

func (r *DeckRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM decks WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
