# Memory Map

A Go-based spaced repetition learning platform that uses the **SM-2 algorithm** to optimize studying and long-term retention. Create decks of flashcards and study efficiently based on scientifically-proven spaced repetition principles.

## Features

- **Spaced Repetition Learning**: Uses the SM-2 algorithm to intelligently schedule card reviews based on difficulty and retention
- **Deck Management**: Create, read, update, and delete study decks
- **Flashcard Management**: Organize cards with front (question) and back (answer) content
- **Adaptive Scheduling**: Cards are automatically rescheduled based on your performance
- **User Authentication**: Secure signup and login with JWT tokens
- **Study Sessions**: Track review history and optimize learning intervals
- **Persistent Storage**: PostgreSQL database for cards, decks, users, and review history

## Architecture

Memory Map follows a clean architecture pattern with distinct layers:

```
cmd/server/              → Application entry point
internal/
  ├── core/              → Business logic (services & domain models)
  │   ├── card/          → Card domain and service
  │   ├── deck/          → Deck domain and service
  │   ├── study/         → SM-2 algorithm and study logic
  │   └── user/          → User domain and service
  ├── platform/          → Infrastructure & frameworks
  │   ├── http/          → HTTP handlers and middleware
  │   └── postgres/      → PostgreSQL repository implementations
  └── adapters/          → Cross-layer adapters
```

## Getting Started

### Prerequisites

- **Go 1.24+**
- **PostgreSQL 12+**
- Make (optional, for Makefile commands)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/snjwilson/memory-map.git
cd memory-map
```

2. Install dependencies:
```bash
go mod download
go mod tidy
```

3. Set up PostgreSQL:
```bash
# Create a database
createdb memorymap
```

4. Run migrations (create tables):
```bash
# You'll need to run SQL scripts to set up the schema
# Tables: users, decks, cards, reviews
```

5. Run the server:
```bash
make run
# or
go run ./cmd/server/main.go
```

The server will start on `http://localhost:8080` by default.

## API Endpoints

### Authentication

- `POST /signup` - Register a new user
- `POST /login` - Login and receive JWT token

### Decks

- `POST /decks` - Create a new deck
- `GET /decks/{id}` - Get deck details
- `GET /user/decks` - Get all decks for the authenticated user
- `PUT /decks/{id}` - Update deck name/description
- `DELETE /decks/{id}` - Delete a deck

### Cards

- `POST /cards` - Create a new card in a deck
- `GET /decks/{id}/cards` - Get all cards in a deck
- `GET /cards/{id}` - Get card details
- `PUT /cards/{id}` - Edit card front/back text
- `DELETE /cards/{id}` - Delete a card

### Study Sessions

- `GET /study/{deckId}` - Get due cards for study
- `POST /study/review` - Submit a review and get next card scheduling

## The SM-2 Algorithm

Memory Map implements the **SuperMemo 2 (SM-2)** spaced repetition algorithm, which calculates optimal review intervals based on:

- **Ease Factor (EF)**: A measure of how easy a card is to remember (default: 2.5)
- **Interval**: Days until the next review
- **Repetitions**: Number of consecutive successful reviews
- **Quality**: Your response quality (1: Again, 2: Hard, 3: Good, 4: Easy)

The algorithm automatically adjusts difficulty and scheduling to maximize learning efficiency.

## Development

### Running Tests

```bash
make test
```

### Code Formatting

```bash
make fmt
```

### Linting

```bash
make vet
```

### Building

```bash
make build
# Binary will be at bin/memory-map
```

## Project Structure Highlights

### Domain Models

- **User**: User account with authentication
- **Deck**: Container for related cards
- **Card**: A flashcard with front/back content and SM-2 state
- **Review**: Tracks user's performance on a card during study

### Services

Each domain has a service layer that handles business logic:
- `card.Service`: Card CRUD and retrieval
- `deck.Service`: Deck management
- `study.Service`: Study session logic and SM-2 calculations
- `user.Service`: User authentication

### Repositories

Repository interfaces define data access contracts:
- `card.Repository`: Persist and retrieve cards
- `deck.Repository`: Persist and retrieve decks
- `study.Repository`: Store and query reviews
- `user.Repository`: Persist and retrieve users

## Environment Configuration

Update your database connection string in `cmd/server/main.go`:

```go
db, err := sql.Open("pgx", "postgres://user:password@localhost:5432/memorymap?sslmode=disable")
```

## Future Enhancements

- Web and mobile frontend
- Card import/export (CSV, Anki format)
- Collaborative deck sharing
- Analytics and study statistics
- Multiple language support
- Offline sync capabilities

## License

MIT

## Author

Sanjay Wilson
