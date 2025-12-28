package http

import (
	"github.com/snjwilson/memory-map/internal/core/card"
	"github.com/snjwilson/memory-map/internal/core/deck"
	"github.com/snjwilson/memory-map/internal/core/study"
)

type Handler struct {
	deckService  *deck.Service
	cardService  *card.Service
	studyService *study.Service
}

func NewHandler(ds *deck.Service, cs *card.Service, ss *study.Service) *Handler {
	return &Handler{
		deckService:  ds,
		cardService:  cs,
		studyService: ss,
	}
}
