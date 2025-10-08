package deck

import (
	"reflect"
	"testing"
)

func TestDeck(t *testing.T) {
	t.Run("NewDeck function returns a Deck type", func(t *testing.T) {
		got := NewDeck("Cats")
		if reflect.TypeOf(*got) != reflect.TypeOf(Deck{}) {
			t.Errorf("Expected a return type of Deck, but got %v", reflect.TypeOf(got))
		}
	})

	t.Run("NewDeck function returns a new Deck with the input name", func(t *testing.T) {
		got := NewDeck("go")
		expected := Deck{Name: "go"}
		if got.Name != expected.Name {
			t.Errorf("Expected a deck with name %v, but got %v", expected.Name, got.Name)
		}
	})

	t.Run("NewDeck function returns a new Deck with an empty slice of cards", func(t *testing.T) {
		got := NewDeck("dogs")
		if len(got.Cards) != 0 {
			t.Errorf("Expected a slice of zero Cards but got %v", len(got.Cards))
		}
	})
}
