package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type CardCollection struct {
	cards []card
}

func NewCardCollection(cap int) *CardCollection {
	return &CardCollection{make([]card, 0, cap)}
}

func MustCollection(c *CardCollection, err error) *CardCollection {
	if err != nil {
		panic(err)
	}

	return c
}

func (c *CardCollection) addCard(card card) error {
	if _, ok := c.cardExists(card); ok {
		return errors.New(fmt.Sprintf("card %q is duplicated in the collection", card.code))
	}

	c.cards = append(c.cards, card)

	return nil
}

func (c *CardCollection) removeCard(card card) error {
	if c.isEmpty() {
		return errors.New("collection is empty, can't remove more cards")
	}

	i, ok := c.cardExists(card)

	if !ok {
		return errors.New("card doesn't exists in the collection")
	}

	c.cards = append(c.cards[:i], c.cards[i+1:]...)

	return nil
}

// integer return is the position of the found card, -1 otherwise
// boolean return indicates card exists or not
func (c *CardCollection) cardExists(card card) (int, bool) {
	for i, c := range c.cards {
		if c.code == card.code {
			return i, true
		}
	}
	return -1, false
}

func (c *CardCollection) isEmpty() bool {
	return len(c.cards) == 0
}

func (c *CardCollection) length() int {
	return len(c.cards)
}

func (c *CardCollection) chooseRandomCard() card {
	rand.Seed(time.Now().UnixNano())
	i := rand.Int() % len(c.cards)
	return c.cards[i]
}

func (c *CardCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Cards []card `json:"cards"`
	}{
		c.cards,
	})
}
