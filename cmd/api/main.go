package main

import (
	"flag"
	"fmt"
	"github.com/HazemNoor/card-game/pkg/player"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var address string
var redisAddress string
var redisPassword string

var gamePlayer *player.Player

func init() {
	flag.StringVar(&address, "http-address", ":8000", "address to listen to, in the form \"host:port\"")
	flag.StringVar(&redisAddress, "redis-address", "127.0.0.1:6379", "address of redis server, in the form \"host:port\"")
	flag.StringVar(&redisPassword, "redis-password", "", "password of redis server")
	flag.Parse()
}

func main() {
	gamePlayer = player.CreateNewPlayer(redisAddress, redisPassword)
	router := mux.NewRouter()
	router.HandleFunc("/deck", createNewDeckHandler).Methods("PUT")
	router.HandleFunc("/deck/{deckId}", openDeckHandler).Methods("GET")
	router.HandleFunc("/deck/{deckId}/draw", drawCardHandler).Methods("POST")

	router.Use(contentTypeMiddleware)
	router.Use(errorMiddleware)

	fmt.Println(fmt.Sprintf("Starting the server on %s", address))

	log.Fatal(http.ListenAndServe(address, router))
}
