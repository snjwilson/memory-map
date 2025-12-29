package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/snjwilson/memory-map/internal/adapters"
	"github.com/snjwilson/memory-map/internal/core/card"
	"github.com/snjwilson/memory-map/internal/core/deck"
	"github.com/snjwilson/memory-map/internal/core/study"
	httpHandler "github.com/snjwilson/memory-map/internal/platform/http"
	"github.com/snjwilson/memory-map/internal/platform/postgres"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:@localhost:5432/dbname=memorymap?sslmode=disable")
	if err != nil {
		fmt.Printf("Error connecting to DB - %v\n", err.Error())
		os.Exit(1)
	}
	defer db.Close()

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
	fmt.Println("Initializing the server...")
	server.ListenAndServe()
}
