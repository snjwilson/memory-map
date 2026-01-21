package postgres

import (
	"context"
	"database/sql"

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

func (r *DeckRepository) ListByOwner(ctx context.Context, ownerId string) ([]*deck.Deck, error) {
	query := `SELECT id, owner_id, name, description, is_public, card_count, created_at, updated_at FROM decks WHERE owner_id = $1`
	rows, err := r.db.QueryContext(ctx, query, ownerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*deck.Deck{}
	for rows.Next() {
		d := &deck.Deck{}
		err := rows.Scan(&d.ID, &d.OwnerID, &d.Name, &d.Description, &d.IsPublic, &d.CardCount, &d.CreatedAt, &d.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, d)
	}
	return result, nil
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
