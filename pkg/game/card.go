package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type card struct {
	value value
	suit  suit
	code  string
}

func newCard(value value, suit suit) *card {
	return &card{value, suit, fmt.Sprintf("%s%s", value.code, suit.code)}
}

func NewCardFromCode(c string) (*card, error) {
	// it has to be at least 2 characters length
	if len(c) < 2 {
		return nil, errors.New(fmt.Sprintf("invalid card code provided %q", c))
	}

	valueCode := c[:len(c)-1]
	v, err := NewValue(valueCode)
	if err != nil {
		return nil, fmt.Errorf("invalid card code %q: %w", c, err)
	}

	suitCode := c[len(valueCode):]
	s, err := NewSuit(suitCode)
	if err != nil {
		return nil, fmt.Errorf("invalid card code %q: %w", c, err)
	}

	return newCard(*v, *s), nil
}

func (c *card) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Value string `json:"value"`
		Suit  string `json:"suit"`
		Code  string `json:"code"`
	}{
		Value: strings.ToUpper(c.value.name),
		Suit:  strings.ToUpper(c.suit.name),
		Code:  c.code,
	})
}
