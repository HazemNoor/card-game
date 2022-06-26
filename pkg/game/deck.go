package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type Deck struct {
	deckId    uuid.UUID
	shuffled  bool
	remaining int
	cards     *CardCollection
}

func NewDeck(id uuid.UUID, shuffled bool, cards ...string) (*Deck, error) {
	cardsCodes := prepareCardsCodes(shuffled, cards)

	d := &Deck{
		deckId:    id,
		shuffled:  shuffled,
		remaining: 0,
		cards:     NewCardCollection(len(cardsCodes)),
	}

	for _, cardCode := range cardsCodes {
		card, err := NewCardFromCode(cardCode)
		if err != nil {
			return nil, fmt.Errorf("invalid card code %q can't add to the deck: %w", cardCode, err)
		}

		err = d.addCard(*card)
		if err != nil {
			return nil, err
		}
	}

	return d, nil
}

func MustDeck(d *Deck, err error) *Deck {
	if err != nil {
		panic(err)
	}

	return d
}

func (d *Deck) DrawCards(n int) (*CardCollection, error) {
	collection := NewCardCollection(n)

	if d.remaining == 0 {
		return nil, errors.New("deck is empty, no more cards to draw")
	}

	if n > d.remaining {
		return nil, errors.New(fmt.Sprintf("can not draw \"%d\" cards max is \"%d\"", n, d.remaining))
	}

	for i := 0; i < n; i++ {
		card, err := d.drawCard()
		if err != nil {
			return collection, err
		}

		if err = collection.addCard(*card); err != nil {
			return nil, err
		}
	}

	return collection, nil
}

func (d *Deck) GetId() uuid.UUID {
	return d.deckId
}

func (d *Deck) GetShuffled() bool {
	return d.shuffled
}

func (d *Deck) GetRemaining() int {
	return d.remaining
}

func (d *Deck) GetCards() CardCollection {
	return *d.cards
}

func (d *Deck) drawCard() (*card, error) {
	if d.cards.isEmpty() {
		return nil, errors.New("deck is empty, can't draw more cards")
	}

	card := d.cards.chooseRandomCard()

	_ = d.removeCard(card)

	return &card, nil
}

func (d *Deck) addCard(card card) error {
	err := d.cards.addCard(card)
	if err != nil {
		return fmt.Errorf("can't add card to the deck: %w", err)
	}

	d.remaining = d.cards.length()

	return nil
}

func (d *Deck) removeCard(card card) error {
	err := d.cards.removeCard(card)
	if err != nil {
		return fmt.Errorf("can't remove card from to the deck: %w", err)
	}

	d.remaining = d.cards.length()

	return nil
}

func (d *Deck) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		DeckId    uuid.UUID `json:"deck_id"`
		Shuffled  bool      `json:"shuffled"`
		Remaining int       `json:"remaining"`
		Cards     []card    `json:"cards"`
	}{
		DeckId:    d.deckId,
		Shuffled:  d.shuffled,
		Remaining: d.remaining,
		Cards:     d.cards.cards,
	})
}

func (d *Deck) UnmarshalJSON(b []byte) error {
	var mapped map[string]interface{}

	jsonDeckError := errors.New("invalid json text for deck")

	if err := json.Unmarshal(b, &mapped); err != nil {
		return err
	}

	var cardsCodes []string
	var err error

	if v, ok := mapped["deck_id"].(string); ok {
		d.deckId, err = uuid.Parse(v)
		if err != nil {
			return err
		}
	} else {
		return jsonDeckError
	}

	if v, ok := mapped["shuffled"].(bool); ok {
		d.shuffled = v
	} else {
		return jsonDeckError
	}

	if v, ok := mapped["cards"].([]interface{}); ok {
		for _, v := range v {
			if v, ok := v.(map[string]interface{}); ok {
				if v, ok := v["code"].(string); ok {
					cardsCodes = append(cardsCodes, v)
				} else {
					return jsonDeckError
				}
			} else {
				return jsonDeckError
			}
		}
	} else {
		return jsonDeckError
	}

	d.cards = NewCardCollection(len(cardsCodes))

	for _, cardCode := range cardsCodes {
		newCard, err := NewCardFromCode(cardCode)
		if err != nil {
			return jsonDeckError
		}
		err = d.addCard(*newCard)
		if err != nil {
			return jsonDeckError
		}
	}

	return nil
}

// filter and shuffle cards
func prepareCardsCodes(shuffled bool, filteredCards []string) []string {
	orderedCards := sequentialOrder()

	var newCards []string

	if len(filteredCards) == 0 {
		newCards = orderedCards[:]
	} else {
		for _, orderedCard := range orderedCards {
			for i, filteredCard := range filteredCards {
				if orderedCard == filteredCard {
					newCards = append(newCards, filteredCard)
					filteredCards[i] = ""
				}
			}
		}
	}

	if len(newCards) != len(filteredCards) {
		for _, filteredCard := range filteredCards {
			if filteredCard != "" {
				newCards = append(newCards, filteredCard)
			}
		}
	}

	if shuffled {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(newCards), func(i, j int) { newCards[i], newCards[j] = newCards[j], newCards[i] })
	}

	return newCards
}

func sequentialOrder() [52]string {
	return [52]string{
		"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
		"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
		"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
		"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
	}
}
