package http

import (
	"encoding/json"
	"net/http"

	"github.com/snjwilson/memory-map/internal/core/deck"
	"github.com/snjwilson/memory-map/internal/platform/http/middleware"
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

func (h *Handler) GetDeckById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing deck id", http.StatusBadRequest)
		return
	}

	deck, err := h.deckService.GetDeckById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deck)
}

func (h *Handler) GetUserDecks(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(middleware.UserIDKey)
	userID, ok := userIDVal.(string)

	if !ok {
		http.Error(w, "missing user id in context", http.StatusBadRequest)
		return
	}

	decks, err := h.deckService.GetUserDecks(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(decks)
}

func (h *Handler) UpdateDeck(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing deck id", http.StatusBadRequest)
		return
	}

	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.deckService.UpdateDeck(r.Context(), id, req.Name, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteDeck(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing deck id", http.StatusBadRequest)
		return
	}

	err := h.deckService.DeleteDeck(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
