package app

import (
	"errors"
	"github.com/HazemNoor/card-game/pkg/game"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	store map[string]game.Deck
}

func NewMemoryRepository() Repository {
	return &MemoryRepository{make(map[string]game.Deck)}
}

func (r MemoryRepository) GetDeck(deckId uuid.UUID) (*game.Deck, error) {
	if deck, ok := r.store[deckId.String()]; ok {
		return &deck, nil
	}
	return nil, errors.New("deck not found")
}

func (r MemoryRepository) SaveDeck(deck *game.Deck) error {
	r.store[deck.GetId().String()] = *deck
	return nil
}
