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
	"github.com/snjwilson/memory-map/internal/core/user"
	httpHandler "github.com/snjwilson/memory-map/internal/platform/http"
	"github.com/snjwilson/memory-map/internal/platform/http/middleware"
	"github.com/snjwilson/memory-map/internal/platform/postgres"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:@localhost:5432/memorymap?sslmode=disable")
	if err != nil {
		fmt.Printf("Error connecting to DB - %v\n", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Wire up the infrastructure
	deckRepo := postgres.NewDeckRepository(db)
	cardRepo := postgres.NewCardRepository(db)
	reviewRepo := postgres.NewReviewRepository(db)
	userRepo := postgres.NewUserRepository(db)

	// Inject into services
	deckService := deck.NewService(deckRepo)
	deckAdapter := adapters.DeckServiceAdapter{Service: deckService}
	cardService := card.NewService(cardRepo, &deckAdapter)
	studyService := study.NewService(reviewRepo, cardRepo)
	userService := user.NewService(userRepo)

	h := httpHandler.NewHandler(deckService, cardService, studyService, userService)
	mux := http.NewServeMux()

	// Apply CORS middleware to all routes
	corsHandler := middleware.CORS(mux)

	// Public Routes
	mux.HandleFunc("POST /signup", h.HandleSignUp)
	mux.HandleFunc("POST /login", h.HandleLogin)

	// Protected Routes

	// Deck Routes
	mux.Handle("GET /decks/:id", middleware.Auth(http.HandlerFunc(h.GetDeckById)))
	mux.Handle("GET /decks/user", middleware.Auth(http.HandlerFunc(h.GetUserDecks)))
	mux.Handle("POST /decks", middleware.Auth(http.HandlerFunc(h.CreateDeck)))
	mux.Handle("PUT /decks", middleware.Auth(http.HandlerFunc(h.UpdateDeck)))
	mux.Handle("DELETE /decks/:id", middleware.Auth(http.HandlerFunc(h.DeleteDeck)))

	// Card Routes
	mux.Handle("POST /cards", middleware.Auth(http.HandlerFunc(h.CreateCard)))
	mux.Handle("GET /decks/:id/cards", middleware.Auth(http.HandlerFunc(h.GetDeckCards)))
	mux.Handle("GET /cards/:id", middleware.Auth(http.HandlerFunc(h.GetCard)))
	mux.Handle("PUT /cards", middleware.Auth(http.HandlerFunc(h.UpdateCard)))
	mux.Handle("DELETE /cards/:id", middleware.Auth(http.HandlerFunc(h.DeleteCard)))

	// Study Routes
	mux.Handle("GET /study/due", middleware.Auth(http.HandlerFunc(h.GetDueCards)))
	mux.Handle("POST study/review", middleware.Auth(http.HandlerFunc(h.SubmitReview)))

	server := &http.Server{
		Addr:    ":8080",
		Handler: corsHandler,
	}
	fmt.Println("Initializing the server...")
	server.ListenAndServe()
}
