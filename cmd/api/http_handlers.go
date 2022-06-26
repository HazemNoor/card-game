package main

import (
	"encoding/json"
	"fmt"
	"github.com/HazemNoor/card-game/pkg/game"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func createNewDeckHandler(w http.ResponseWriter, r *http.Request) {
	var shuffled bool
	var cards []string

	if r.URL.Query().Get("shuffled") == "1" {
		shuffled = true
	}

	cards = strings.FieldsFunc(r.URL.Query().Get("cards"), func(c rune) bool { return c == ',' })

	deck := game.MustDeck(gameApp.CreateNewDeck(shuffled, cards))

	payload := struct {
		DeckId    string `json:"deck_id"`
		Shuffled  bool   `json:"shuffled"`
		Remaining int    `json:"remaining"`
	}{
		DeckId:    deck.GetId().String(),
		Shuffled:  deck.GetShuffled(),
		Remaining: deck.GetRemaining(),
	}

	w.WriteHeader(http.StatusCreated)
	must(json.NewEncoder(w).Encode(payload))
}

func openDeckHandler(w http.ResponseWriter, r *http.Request) {
	deckId := mux.Vars(r)["deckId"]

	payload := game.MustDeck(gameApp.OpenDeck(deckId))

	w.WriteHeader(http.StatusOK)
	must(json.NewEncoder(w).Encode(payload))
}

func drawCardHandler(w http.ResponseWriter, r *http.Request) {
	deckId := mux.Vars(r)["deckId"]

	var count int
	_, _ = fmt.Sscan(r.URL.Query().Get("count"), &count)
	if count == 0 {
		count = 1
	}

	payload := game.MustCollection(gameApp.DrawCard(deckId, count))

	w.WriteHeader(http.StatusOK)
	must(json.NewEncoder(w).Encode(payload))
}

func contentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func errorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusBadRequest)

				payload := make(map[string]string)
				payload["success"] = "false"
				payload["message"] = fmt.Sprintf("%v", err)

				_ = json.NewEncoder(w).Encode(payload)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
