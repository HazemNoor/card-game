package app

import (
	"github.com/HazemNoor/card-game/pkg/game"
	"github.com/google/uuid"
)

type Repository interface {
	GetDeck(deckId uuid.UUID) (*game.Deck, error)
	SaveDeck(deck *game.Deck) error
}
