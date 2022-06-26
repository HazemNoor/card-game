package game

import (
	"encoding/json"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestDeck_DrawCards(t *testing.T) {
	deck := MustDeck(NewDeck(uuid.New(), false))

	if deck.remaining != 52 {
		t.Errorf("Deck.remaining got = %v, want %v", deck.remaining, 52)
	}

	_, _ = deck.DrawCards(3)

	if deck.remaining != 49 {
		t.Errorf("Deck.remaining got = %v, want %v", deck.remaining, 49)
	}
}

func TestDeck_DrawTooManyCards(t *testing.T) {
	deck := MustDeck(NewDeck(uuid.New(), false))

	_, err := deck.DrawCards(60)

	if err == nil {
		t.Errorf("should not draw cards greater than what is remaining in the deck")
	}
}

func TestNewOrderedDeck(t *testing.T) {
	id := uuid.New()
	want := &Deck{
		deckId:    id,
		shuffled:  false,
		remaining: 3,
		cards: &CardCollection{[]card{
			{value{"A", "Ace"}, suit{"S", "Spades"}, "AS"},
			{value{"2", "2"}, suit{"S", "Spades"}, "2S"},
			{value{"3", "3"}, suit{"S", "Spades"}, "3S"},
		}},
	}

	got, _ := NewDeck(id, false, "AS", "2S", "3S")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewDeck() got = %v, want %v", got, want)
	}
}

func TestNewShuffledDeck(t *testing.T) {
	id := uuid.New()
	orderedDeck := &Deck{
		deckId:    id,
		shuffled:  false,
		remaining: 10,
		cards: &CardCollection{[]card{
			{value{"A", "Ace"}, suit{"S", "Spades"}, "AS"},
			{value{"2", "2"}, suit{"S", "Spades"}, "2S"},
			{value{"3", "3"}, suit{"S", "Spades"}, "3S"},
			{value{"4", "4"}, suit{"S", "Spades"}, "4S"},
			{value{"5", "5"}, suit{"S", "Spades"}, "5S"},
			{value{"6", "6"}, suit{"S", "Spades"}, "6S"},
			{value{"7", "7"}, suit{"S", "Spades"}, "7S"},
			{value{"8", "8"}, suit{"S", "Spades"}, "8S"},
			{value{"9", "9"}, suit{"S", "Spades"}, "9S"},
			{value{"10", "10"}, suit{"S", "Spades"}, "10S"},
		}},
	}

	shuffledDeck, _ := NewDeck(id, true, "AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S")
	if reflect.DeepEqual(shuffledDeck, orderedDeck) {
		t.Errorf("new Deck is sorted, must be shuffled instead")
	}

	if shuffledDeck.remaining != 10 {
		t.Errorf("new Deck remaining = %d, want %d", shuffledDeck.remaining, 10)
	}
}

func TestCannotCreateNewDeckWithDuplicatedCards(t *testing.T) {
	_, err := NewDeck(uuid.New(), false, "AS", "AS")
	if err == nil {
		t.Errorf("NewDeck() should not create deck with duplicated cards")
	}
}

func TestCannotCreateNewDeckWithInvalidCards(t *testing.T) {
	_, err := NewDeck(uuid.New(), false, "AS", "AZ")
	if err == nil {
		t.Errorf("NewDeck() should not create deck with invalid cards")
	}
}

func TestDeckJsonMarshalAndUnMarshal(t *testing.T) {
	deck := MustDeck(NewDeck(uuid.New(), false, "AS", "2S", "3S"))

	jsonData, _ := json.Marshal(deck)

	decodedDeck := &Deck{}
	_ = json.Unmarshal(jsonData, decodedDeck)

	if !reflect.DeepEqual(deck, decodedDeck) {
		t.Errorf("json.Marshal, json.Unmarshal error for Deck")
	}
}

func Test_sequentialOrder(t *testing.T) {
	want := [52]string{
		"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
		"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
		"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
		"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
	}
	if got := sequentialOrder(); !reflect.DeepEqual(got, want) {
		t.Errorf("sequentialOrder() = %v, want %v", got, want)
	}
}
