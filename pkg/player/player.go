package player

import (
	"github.com/HazemNoor/card-game/pkg/game"
	"github.com/google/uuid"
)

type Player struct {
	repo Repository
}

func NewPlayer(repository Repository) *Player {
	return &Player{repo: repository}
}

func (p *Player) CreateNewDeck(shuffled bool, cards []string) (*game.Deck, error) {
	deck, err := game.NewDeck(uuid.New(), shuffled, cards...)
	if err != nil {
		return nil, err
	}

	err = p.repo.SaveDeck(deck)
	if err != nil {
		return nil, err
	}

	return deck, nil
}

func (p *Player) OpenDeck(deckId string) (*game.Deck, error) {
	deckUUID, err := uuid.Parse(deckId)
	if err != nil {
		return nil, err
	}

	deck, err := p.repo.GetDeck(deckUUID)
	if err != nil {
		return nil, err
	}

	return deck, nil
}

func (p *Player) DrawCard(deckId string, n int) (*game.CardCollection, error) {
	deck, err := p.OpenDeck(deckId)
	if err != nil {
		return nil, err
	}

	collection, err := deck.DrawCards(n)
	if err != nil {
		return nil, err
	}

	err = p.repo.SaveDeck(deck)
	if err != nil {
		return nil, err
	}

	return collection, nil
}
