# Another Card Game

An implementation of a REST API to simulate a deck of cards.

### Credit
Hazem Noor <hazemnoor@gmail.com>

## How to run this code

1. At least `go1.18` and `docker` needs to be installed
2. You can see list of commands using make command
```shell
make
```
3. Run tests
```shell
make test
```
4. Start the Redis server by docker
```shell
make start-redis
```
it will run on port `6379` by default, you can change from `.env` file

5. Build and Start the binary
```shell
make start
```
it will run on port `8000` by default, you can change from `.env` file

the binary file will be located in `bin` directory

---

To stop the server hit `Ctrl+C`

Don't forget to stop redis too with `make stop-redis`

---

## API Documentation
- Create a new Deck
```shell
curl --request PUT --url 'http://localhost:8000/deck?shuffled=1&cards=AS,KD,AC'
```
Response would be
```json
{
    "deck_id": "05d5b426-a086-48d5-ad9c-1c913774154b",
    "shuffled": true,
    "remaining": 3
}
```

- Open a Deck
```shell
curl --request GET --url 'http://localhost:8000/deck/05d5b426-a086-48d5-ad9c-1c913774154b'
```
Response would be
```json
{
    "deck_id": "05d5b426-a086-48d5-ad9c-1c913774154b",
    "shuffled": true,
    "remaining": 3,
    "cards": [
        {
            "value": "KING",
            "suit": "DIAMONDS",
            "code": "KD"
        },
        {
            "value": "ACE",
            "suit": "SPADES",
            "code": "AS"
        },
        {
            "value": "ACE",
            "suit": "CLUBS",
            "code": "AC"
        }
    ]
}
```

- Draw a Card
```shell
curl --request POST --url 'http://localhost:8000/deck/05d5b426-a086-48d5-ad9c-1c913774154b/draw?count=2'
```
Response would be
```json
{
    "cards": [
        {
            "value": "ACE",
            "suit": "SPADES",
            "code": "AS"
        },
        {
            "value": "KING",
            "suit": "DIAMONDS",
            "code": "KD"
        }
    ]
}
```
