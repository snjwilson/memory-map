package http

import (
	"encoding/json"
	"net/http"

	"github.com/snjwilson/memory-map/internal/core/card"
)

// CreateCard handles POST /cards
func (h *Handler) CreateCard(w http.ResponseWriter, r *http.Request) {
	var req struct {
		DeckID string `json:"deck_id"`
		Front  string `json:"front"`
		Back   string `json:"back"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	c, err := h.cardService.CreateCard(r.Context(), card.NewCardRequest{
		DeckID: req.DeckID,
		Front:  req.Front,
		Back:   req.Back,
	})
	if err != nil {
		if err == card.ErrInvalidContent {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(c)
}

// GetDeckCards handles GET /decks/{id}/cards
func (h *Handler) GetDeckCards(w http.ResponseWriter, r *http.Request) {
	deckID := r.PathValue("id")
	if deckID == "" {
		http.Error(w, "missing deck id", http.StatusBadRequest)
		return
	}

	cards, err := h.cardService.GetByDeckId(r.Context(), deckID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(cards)
}

// GetCard handles GET /cards/{id}
func (h *Handler) GetCard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing card id", http.StatusBadRequest)
		return
	}

	c, err := h.cardService.GetById(r.Context(), id)
	if err != nil {
		if err == card.ErrCardNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(c)
}

// UpdateCard handles PUT /cards/{id}
func (h *Handler) UpdateCard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing card id", http.StatusBadRequest)
		return
	}

	var req struct {
		Front string `json:"front"`
		Back  string `json:"back"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.cardService.UpdateCard(r.Context(), id, req.Front, req.Back); err != nil {
		if err == card.ErrInvalidContent {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err == card.ErrCardNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteCard handles DELETE /cards/{id}
func (h *Handler) DeleteCard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing card id", http.StatusBadRequest)
		return
	}

	if err := h.cardService.DeleteCard(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
