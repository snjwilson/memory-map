package main

import (
	"database/sql"
	"net/http"

	"github.com/snjwilson/memory-map/internal/adapters"
	"github.com/snjwilson/memory-map/internal/core/card"
	"github.com/snjwilson/memory-map/internal/core/deck"
	"github.com/snjwilson/memory-map/internal/core/study"
	httpHandler "github.com/snjwilson/memory-map/internal/platform/http"
	"github.com/snjwilson/memory-map/internal/platform/postgres"
)

func main() {
	db, _ := sql.Open("postgres", "postgres://user:pass@localhost/dbname?sslmode=disable")

	// Wire up the infrastructure
	deckRepo := postgres.NewDeckRepository(db)
	cardRepo := postgres.NewCardRepository(db)
	reviewRepo := postgres.NewReviewRepository(db)

	// Inject into services
	deckService := deck.NewService(deckRepo)
	deckAdapter := adapters.DeckServiceAdapter{Service: deckService}
	cardService := card.NewService(cardRepo, &deckAdapter)
	studyService := study.NewService(reviewRepo, cardRepo)

	h := httpHandler.NewHandler(deckService, cardService, studyService)
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("POST /decks", h.CreateDeck)
	mux.HandleFunc("GET /cards/due", h.GetDueCards)
	mux.HandleFunc("POST /reviews", h.SubmitReview)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
