package game

import (
	"errors"
	"fmt"
	"strings"
)

type value struct {
	code, name string
}

func NewValue(c string) (*value, error) {
	c = strings.ToUpper(c)
	if n, ok := allValues()[c]; ok {
		return &value{c, n}, nil
	}

	return nil, errors.New(fmt.Sprintf("invalid value code %q", c))
}

func allValues() map[string]string {
	return map[string]string{
		"A":  "Ace",
		"2":  "2",
		"3":  "3",
		"4":  "4",
		"5":  "5",
		"6":  "6",
		"7":  "7",
		"8":  "8",
		"9":  "9",
		"10": "10",
		"J":  "Jack",
		"Q":  "Queen",
		"K":  "King",
	}
}
