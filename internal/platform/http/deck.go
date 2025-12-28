package http

import (
	"encoding/json"
	"net/http"

	"github.com/snjwilson/memory-map/internal/core/deck"
)

func (h *Handler) CreateDeck(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		OwnerID     string `json:"owner_id"` // In production, get this from JWT context
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	deck, err := h.deckService.CreateDeck(r.Context(), deck.NewDeckRequest{
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     req.OwnerID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deck)
}
