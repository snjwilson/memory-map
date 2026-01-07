package http

import (
	"encoding/json"
	"net/http"

	"github.com/snjwilson/memory-map/internal/core/study"
)

// GetDueCards fetches the cards the user needs to study right now
func (h *Handler) GetDueCards(w http.ResponseWriter, r *http.Request) {
	deckID := r.PathValue("deckId")
	if deckID == "" {
		http.Error(w, "deck id is required", http.StatusBadRequest)
		return
	}

	cards, err := h.cardService.GetDueCards(r.Context(), deckID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}

// SubmitReview processes the answer
func (h *Handler) SubmitReview(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CardID     string       `json:"cardId"`
		Rating     study.Rating `json:"rating"` // 1-4
		DurationMs int          `json:"durationMs"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	updatedCard, err := h.studyService.SubmitReview(r.Context(), req.CardID, req.Rating, req.DurationMs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the updated card so the UI knows the next DueDate
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCard)
}
