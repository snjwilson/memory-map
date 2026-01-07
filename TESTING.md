# Memory Map API Testing Guide

Complete request/response payloads for testing all endpoints in Postman or any HTTP client.

---

## Authentication Endpoints

### 1. Sign Up
**Endpoint:** `POST /signup`  
**Auth:** None (Public)

**Request:**
```json
{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "user@example.com",
  "created_at": "2026-01-03T10:30:00Z",
  "updated_at": "2026-01-03T10:30:00Z"
}
```

---

### 2. Login
**Endpoint:** `POST /login`  
**Auth:** None (Public)

**Request:**
```json
{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

**Response (200 OK):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "user@example.com",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "created_at": "2026-01-03T10:30:00Z",
  "updated_at": "2026-01-03T10:30:00Z"
}
```

> **Note:** Use the `token` in the `Authorization: Bearer <token>` header for all protected routes.

---

## Deck Endpoints

### 3. Create Deck
**Endpoint:** `POST /decks`  
**Auth:** Required (Bearer Token)

**Request:**
```json
{
  "name": "Spanish Vocabulary",
  "description": "Learn common Spanish words and phrases",
  "owner_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

**Response (201 Created):**
```json
{
  "id": "660e8400-e29b-41d4-a716-446655440001",
  "owner_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Spanish Vocabulary",
  "description": "Learn common Spanish words and phrases",
  "created_at": "2026-01-03T11:00:00Z",
  "updated_at": "2026-01-03T11:00:00Z"
}
```

---

### 4. Get Deck by ID
**Endpoint:** `GET /decks/{id}`  
**Auth:** Required (Bearer Token)

**Example URL:**
```
GET /decks/660e8400-e29b-41d4-a716-446655440001
```

**Response (200 OK):**
```json
{
  "id": "660e8400-e29b-41d4-a716-446655440001",
  "owner_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Spanish Vocabulary",
  "description": "Learn common Spanish words and phrases",
  "created_at": "2026-01-03T11:00:00Z",
  "updated_at": "2026-01-03T11:00:00Z"
}
```

---

### 5. Get User Decks
**Endpoint:** `GET /decks/user`  
**Auth:** Required (Bearer Token)

**Response (200 OK):**
```json
[
  {
    "id": "660e8400-e29b-41d4-a716-446655440001",
    "owner_id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Spanish Vocabulary",
    "description": "Learn common Spanish words and phrases",
    "created_at": "2026-01-03T11:00:00Z",
    "updated_at": "2026-01-03T11:00:00Z"
  },
  {
    "id": "770e8400-e29b-41d4-a716-446655440002",
    "owner_id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "French Grammar",
    "description": "Essential French grammar rules",
    "created_at": "2026-01-02T14:30:00Z",
    "updated_at": "2026-01-02T14:30:00Z"
  }
]
```

---

### 6. Update Deck
**Endpoint:** `PUT /decks/{id}`  
**Auth:** Required (Bearer Token)

**Example URL:**
```
PUT /decks/660e8400-e29b-41d4-a716-446655440001
```

**Request:**
```json
{
  "name": "Spanish Vocabulary - Intermediate",
  "description": "Learn intermediate Spanish words and phrases"
}
```

**Response (204 No Content)**

---

### 7. Delete Deck
**Endpoint:** `DELETE /decks/{id}`  
**Auth:** Required (Bearer Token)

**Example URL:**
```
DELETE /decks/660e8400-e29b-41d4-a716-446655440001
```

**Response (204 No Content)**

---

## Card Endpoints

### 8. Create Card
**Endpoint:** `POST /cards`  
**Auth:** Required (Bearer Token)

**Request:**
```json
{
  "deck_id": "660e8400-e29b-41d4-a716-446655440001",
  "front": "What is the Spanish word for 'hello'?",
  "back": "Hola"
}
```

**Response (201 Created):**
```json
{
  "id": "880e8400-e29b-41d4-a716-446655440003",
  "deck_id": "660e8400-e29b-41d4-a716-446655440001",
  "front": "What is the Spanish word for 'hello'?",
  "back": "Hola",
  "interval": 0,
  "ease_factor": 2.5,
  "repetitions": 0,
  "due_date": "2026-01-03T11:15:00Z",
  "created_at": "2026-01-03T11:15:00Z",
  "updated_at": "2026-01-03T11:15:00Z"
}
```

---

### 9. Get All Cards in a Deck
**Endpoint:** `GET /decks/{id}/cards`  
**Auth:** Required (Bearer Token)

**Example URL:**
```
GET /decks/660e8400-e29b-41d4-a716-446655440001/cards
```

**Response (200 OK):**
```json
[
  {
    "id": "880e8400-e29b-41d4-a716-446655440003",
    "deck_id": "660e8400-e29b-41d4-a716-446655440001",
    "front": "What is the Spanish word for 'hello'?",
    "back": "Hola",
    "interval": 0,
    "ease_factor": 2.5,
    "repetitions": 0,
    "due_date": "2026-01-03T11:15:00Z",
    "created_at": "2026-01-03T11:15:00Z",
    "updated_at": "2026-01-03T11:15:00Z"
  },
  {
    "id": "990e8400-e29b-41d4-a716-446655440004",
    "deck_id": "660e8400-e29b-41d4-a716-446655440001",
    "front": "What is the Spanish word for 'goodbye'?",
    "back": "Adiós",
    "interval": 1,
    "ease_factor": 2.6,
    "repetitions": 1,
    "due_date": "2026-01-04T11:15:00Z",
    "created_at": "2026-01-03T11:20:00Z",
    "updated_at": "2026-01-03T11:20:00Z"
  }
]
```

---

### 10. Get Single Card
**Endpoint:** `GET /cards/{id}`  
**Auth:** Required (Bearer Token)

**Example URL:**
```
GET /cards/880e8400-e29b-41d4-a716-446655440003
```

**Response (200 OK):**
```json
{
  "id": "880e8400-e29b-41d4-a716-446655440003",
  "deck_id": "660e8400-e29b-41d4-a716-446655440001",
  "front": "What is the Spanish word for 'hello'?",
  "back": "Hola",
  "interval": 0,
  "ease_factor": 2.5,
  "repetitions": 0,
  "due_date": "2026-01-03T11:15:00Z",
  "created_at": "2026-01-03T11:15:00Z",
  "updated_at": "2026-01-03T11:15:00Z"
}
```

---

### 11. Update Card
**Endpoint:** `PUT /cards/{id}`  
**Auth:** Required (Bearer Token)

**Example URL:**
```
PUT /cards/880e8400-e29b-41d4-a716-446655440003
```

**Request:**
```json
{
  "front": "What is the Spanish word for 'hello' (informal)?",
  "back": "Hola / Hola, ¿qué tal?"
}
```

**Response (204 No Content)**

---

### 12. Delete Card
**Endpoint:** `DELETE /cards/{id}`  
**Auth:** Required (Bearer Token)

**Example URL:**
```
DELETE /cards/880e8400-e29b-41d4-a716-446655440003
```

**Response (204 No Content)**

---

## Study Endpoints

### 13. Get Due Cards (For Study Session)
**Endpoint:** `GET /study/due`  
**Auth:** Required (Bearer Token)

**Query Parameters:**
- `deck_id` (required): UUID of the deck to study
- `limit` (optional): Number of cards to fetch (default: 20)

**Example URL:**
```
GET /study/due?deck_id=660e8400-e29b-41d4-a716-446655440001&limit=10
```

**Response (200 OK):**
```json
[
  {
    "id": "880e8400-e29b-41d4-a716-446655440003",
    "deck_id": "660e8400-e29b-41d4-a716-446655440001",
    "front": "What is the Spanish word for 'hello'?",
    "back": "Hola",
    "interval": 0,
    "ease_factor": 2.5,
    "repetitions": 0,
    "due_date": "2026-01-03T11:15:00Z",
    "created_at": "2026-01-03T11:15:00Z",
    "updated_at": "2026-01-03T11:15:00Z"
  }
]
```

---

### 14. Submit Review
**Endpoint:** `POST /study/review`  
**Auth:** Required (Bearer Token)

**Request:**
```json
{
  "card_id": "880e8400-e29b-41d4-a716-446655440003",
  "rating": 3,
  "duration_ms": 2500
}
```

**Rating Scale:**
- `1` = Again (complete blackout, need to see immediately)
- `2` = Hard (remembered but with difficulty)
- `3` = Good (correct response with some hesitation)
- `4` = Easy (perfect recall)

**Response (200 OK):**
```json
{
  "id": "aa0e8400-e29b-41d4-a716-446655440005",
  "card_id": "880e8400-e29b-41d4-a716-446655440003",
  "rating": 3,
  "review_time": "2026-01-03T11:30:00Z",
  "duration_ms": 2500,
  "new_interval": 6,
  "new_ease": 2.5
}
```

> **Note:** The SM-2 algorithm automatically calculates the next review interval and ease factor based on your rating.

---

## Testing Workflow

### Complete Study Session Example

1. **Sign up / Login** → Get JWT token
2. **Create a deck** → Get deck ID
3. **Create multiple cards** → Get card IDs
4. **Get due cards** → See which cards need review
5. **Submit reviews** → Rate cards and update SM-2 state
6. **Track progress** → See intervals increase for good recalls

---

## Error Response Examples

### 400 Bad Request (Missing field)
```json
{
  "error": "invalid request body"
}
```

### 401 Unauthorized (Invalid/missing token)
```json
{
  "error": "missing or invalid authorization token"
}
```

### 404 Not Found (Resource doesn't exist)
```json
{
  "error": "card not found"
}
```

### 409 Conflict (Email already taken)
```json
{
  "error": "email already in use"
}
```

### 500 Internal Server Error
```json
{
  "error": "internal server error"
}
```

---

## Quick Test Scenarios

### Scenario 1: New User Learning Spanish
```bash
# 1. Sign up
POST /signup
{ "email": "student@example.com", "password": "pass123" }

# 2. Create Spanish deck
POST /decks
{ "name": "Spanish 101", "description": "Beginner Spanish" }

# 3. Add cards
POST /cards (5 times with different words)
{ "deck_id": "...", "front": "Spanish for X", "back": "Y" }

# 4. Start studying
GET /study/due?deck_id=...

# 5. Rate a card as "Good"
POST /study/review
{ "card_id": "...", "rating": 3, "duration_ms": 3000 }
```

### Scenario 2: Review Interval Progression
- Card rated "Good" (rating 3) → New interval: 6 days
- Same card rated "Easy" (rating 4) after 6 days → New interval: ~15 days
- Same card rated "Hard" (rating 2) after 15 days → Reset to review soon

---

## Postman Collection Tips

- Store `base_url` as: `http://localhost:8080`
- Store `auth_token` from login response
- Use `{{base_url}}` and `{{auth_token}}` in requests
- Add `Authorization: Bearer {{auth_token}}` to protected route headers
- Test complete workflows with Postman's collection runner
