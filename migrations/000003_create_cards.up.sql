-- 1. Create Cards table
CREATE TABLE cards (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    deck_id UUID NOT NULL REFERENCES decks(id) ON DELETE CASCADE,
    front TEXT NOT NULL,
    back TEXT NOT NULL,
    
    -- SM-2 Algorithm State
    interval INTEGER NOT NULL DEFAULT 0,
    -- DOUBLE PRECISION matches Go's float64
    ease_factor DOUBLE PRECISION NOT NULL DEFAULT 2.5 CHECK (ease_factor >= 1.3),
    repetitions INTEGER NOT NULL DEFAULT 0,
    due_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. Performance Indices
-- Optimization for: "Get all cards for this deck"
CREATE INDEX idx_cards_deck_id ON cards(deck_id);
-- Optimization for: "Get all cards in this deck that are due for review"
CREATE INDEX idx_cards_deck_due ON cards(deck_id, due_date);