package deck

import "github.com/google/uuid"

type Deck struct {
	Id    uuid.UUID
	Name  string
	Cards []Card
}

type Card struct {
	Id       uuid.UUID
	Question string
	Answer   string
}

func NewDeck(name string) *Deck {
	return &Deck{
		Id:    uuid.New(),
		Name:  name,
		Cards: []Card{},
	}
}
