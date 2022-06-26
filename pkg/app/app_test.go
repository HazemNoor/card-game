package app

import (
	"github.com/HazemNoor/card-game/pkg/game"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestApp_CreateNewDeck(t *testing.T) {
	testApp := CreateNewAppInMemory()

	_, err := testApp.CreateNewDeck(false, []string{"99"})
	if err == nil {
		t.Error("App.CreateNewDeck() should return an error")
	}

	got, err := testApp.CreateNewDeck(true, []string{"AS"})
	if got == nil {
		t.Fatal("App.CreateNewDeck() returned wrong deck")
	}
	want := game.MustDeck(game.NewDeck(uuid.New(), true, "AS"))

	shuffledEqual := got.GetShuffled() == want.GetShuffled()
	remainingEqual := got.GetRemaining() == want.GetRemaining()
	cardsEqual := reflect.DeepEqual(got.GetCards(), want.GetCards())

	if !shuffledEqual || !remainingEqual || !cardsEqual {
		t.Errorf("App.CreateNewDeck() returned wrong deck")
	}
}

func TestApp_DrawCard(t *testing.T) {
	testApp := CreateNewAppInMemory()
	deck, _ := testApp.CreateNewDeck(true, []string{"AS"})

	gotCollection, err := deck.DrawCards(1)
	if err != nil || gotCollection == nil {
		t.Fatal("DrawCards() returned wrong collection")
	}

	wantDeck := game.MustDeck(game.NewDeck(uuid.New(), true, "AS"))

	wantCollection := game.MustCollection(wantDeck.DrawCards(1))

	if !reflect.DeepEqual(gotCollection, wantCollection) {
		t.Error("App.DrawCard() returned wrong deck")
	}
}

func TestApp_OpenDeck(t *testing.T) {
	testApp := CreateNewAppInMemory()

	want, _ := testApp.CreateNewDeck(false, []string{"AS"})

	got, _ := testApp.OpenDeck(want.GetId().String())

	if !reflect.DeepEqual(got, want) {
		t.Error("OpenDeck() returned wrong deck")
	}
}
