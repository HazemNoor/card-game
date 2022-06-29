package player

import (
	"github.com/HazemNoor/card-game/pkg/game"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestPlayer_CreateNewDeck(t *testing.T) {
	testPlayer := CreateNewPlayerInMemory()

	_, err := testPlayer.CreateNewDeck(false, []string{"99"})
	if err == nil {
		t.Error("Player.CreateNewDeck() should return an error")
	}

	got, err := testPlayer.CreateNewDeck(true, []string{"AS"})
	if got == nil {
		t.Fatal("Player.CreateNewDeck() returned wrong deck")
	}
	want := game.MustDeck(game.NewDeck(uuid.New(), true, "AS"))

	shuffledEqual := got.GetShuffled() == want.GetShuffled()
	remainingEqual := got.GetRemaining() == want.GetRemaining()
	cardsEqual := reflect.DeepEqual(got.GetCards(), want.GetCards())

	if !shuffledEqual || !remainingEqual || !cardsEqual {
		t.Errorf("Player.CreateNewDeck() returned wrong deck")
	}
}

func TestPlayer_DrawCard(t *testing.T) {
	testPlayer := CreateNewPlayerInMemory()
	deck, _ := testPlayer.CreateNewDeck(true, []string{"AS"})

	gotCollection, err := deck.DrawCards(1)
	if err != nil || gotCollection == nil {
		t.Fatal("DrawCards() returned wrong collection")
	}

	wantDeck := game.MustDeck(game.NewDeck(uuid.New(), true, "AS"))

	wantCollection := game.MustCollection(wantDeck.DrawCards(1))

	if !reflect.DeepEqual(gotCollection, wantCollection) {
		t.Error("Player.DrawCard() returned wrong deck")
	}
}

func TestPlayer_OpenDeck(t *testing.T) {
	testPlayer := CreateNewPlayerInMemory()

	want, _ := testPlayer.CreateNewDeck(false, []string{"AS"})

	got, _ := testPlayer.OpenDeck(want.GetId().String())

	if !reflect.DeepEqual(got, want) {
		t.Error("OpenDeck() returned wrong deck")
	}
}
