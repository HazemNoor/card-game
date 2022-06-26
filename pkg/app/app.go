package app

import (
	"github.com/HazemNoor/card-game/pkg/game"
	"github.com/google/uuid"
)

type App struct {
	repo Repository
}

func NewApp(repository Repository) *App {
	return &App{repo: repository}
}

func (app *App) CreateNewDeck(shuffled bool, cards []string) (*game.Deck, error) {
	deck, err := game.NewDeck(uuid.New(), shuffled, cards...)
	if err != nil {
		return nil, err
	}

	err = app.repo.SaveDeck(deck)
	if err != nil {
		return nil, err
	}

	return deck, nil
}

func (app *App) OpenDeck(deckId string) (*game.Deck, error) {
	deckUUID, err := uuid.Parse(deckId)
	if err != nil {
		return nil, err
	}

	deck, err := app.repo.GetDeck(deckUUID)
	if err != nil {
		return nil, err
	}

	return deck, nil
}

func (app *App) DrawCard(deckId string, n int) (*game.CardCollection, error) {
	deck, err := app.OpenDeck(deckId)
	if err != nil {
		return nil, err
	}

	collection, err := deck.DrawCards(n)
	if err != nil {
		return nil, err
	}

	err = app.repo.SaveDeck(deck)
	if err != nil {
		return nil, err
	}

	return collection, nil
}
