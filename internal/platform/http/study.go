package http

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/snjwilson/memory-map/internal/core/study"
)

// GetDeckDueCards fetches the cards the user needs to study right now
func (h *Handler) GetDeckDueCards(w http.ResponseWriter, r *http.Request) {
	deckID := r.PathValue("deckId")
	if deckID == "" {
		http.Error(w, "deck id is required", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit < 1 || limit > 100 {
		limit = 10 // Default items per page
	}

	cards, err := h.cardService.GetDueCards(r.Context(), deckID, page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}

// SubmitReview processes the answer
func (h *Handler) SubmitReview(w http.ResponseWriter, r *http.Request) {
	var req study.StudySessionReviewRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("JSON decode failure", "error", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Pass the slice of reviews to the service
	updatedCards, err := h.studyService.ProcessReview(r.Context(), req.Reviews)
	if err != nil {
		http.Error(w, "failed to process reviews", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Return the list of updated cards
	json.NewEncoder(w).Encode(updatedCards)
}
