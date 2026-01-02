package http

import (
	"github.com/snjwilson/memory-map/internal/core/card"
	"github.com/snjwilson/memory-map/internal/core/deck"
	"github.com/snjwilson/memory-map/internal/core/study"
	"github.com/snjwilson/memory-map/internal/core/user"
)

type Handler struct {
	deckService  *deck.Service
	cardService  *card.Service
	studyService *study.Service
	userService  *user.Service
}

func NewHandler(ds *deck.Service, cs *card.Service, ss *study.Service, us *user.Service) *Handler {
	return &Handler{
		deckService:  ds,
		cardService:  cs,
		studyService: ss,
		userService:  us,
	}
}
