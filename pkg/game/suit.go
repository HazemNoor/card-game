package game

import (
	"errors"
	"fmt"
	"strings"
)

type suit struct {
	code, name string
}

func NewSuit(c string) (*suit, error) {
	c = strings.ToUpper(c)
	if n, ok := allSuits()[c]; ok {
		return &suit{c, n}, nil
	}

	return nil, errors.New(fmt.Sprintf("invalid suit code %q", c))
}

func allSuits() map[string]string {
	return map[string]string{
		"S": "Spades",
		"D": "Diamonds",
		"C": "Clubs",
		"H": "Hearts",
	}
}
