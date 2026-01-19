-- 1. Create Review Logs table
CREATE TABLE review_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    card_id UUID NOT NULL REFERENCES cards(id) ON DELETE CASCADE,
    
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 3),
    
    review_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    duration_ms INTEGER NOT NULL, -- milliseconds
    
    -- Snapshots for history/debugging
    new_interval INTEGER NOT NULL,
    new_ease DOUBLE PRECISION NOT NULL,

    -- Audit trail (no trigger needed as these are immutable logs)
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. Performance Indices
-- For "Show me the history of this specific card"
CREATE INDEX idx_review_logs_card_id ON review_logs(card_id);
-- For "Show me all reviews from last week" (Analytics/Heatmap)
CREATE INDEX idx_review_logs_review_time ON review_logs(review_time);