-- 1. Create the function that updates the count
CREATE OR REPLACE FUNCTION fn_sync_deck_card_count()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        UPDATE decks 
        SET card_count = card_count + 1 
        WHERE id = NEW.deck_id;
    ELSIF (TG_OP = 'DELETE') THEN
        UPDATE decks 
        SET card_count = card_count - 1 
        WHERE id = OLD.deck_id;
    END IF;
    RETURN NULL; -- result is ignored since this is an AFTER trigger
END;
$$ LANGUAGE plpgsql;

-- 2. Create the trigger to fire after any change to the cards table
CREATE TRIGGER trg_sync_card_count
AFTER INSERT OR DELETE ON cards
FOR EACH ROW
EXECUTE FUNCTION fn_sync_deck_card_count();

-- 3. Initial Sync
-- Just in case there are already cards in the DB, this aligns the counts
UPDATE decks d
SET card_count = (SELECT count(*) FROM cards c WHERE c.deck_id = d.id);