package app

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/HazemNoor/card-game/pkg/game"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
)

type RedisRepository struct {
	ctx context.Context
	db  *redis.Client
}

func NewRedisRepository(addr, pass string) Repository {
	return &RedisRepository{
		ctx: context.Background(),
		db: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pass,
			DB:       0,
		}),
	}
}

func (r RedisRepository) GetDeck(deckId uuid.UUID) (*game.Deck, error) {
	jsonData, err := r.db.Get(r.ctx, deckId.String()).Result()
	if err == redis.Nil {
		return nil, errors.New("deck not found")
	} else if err != nil {
		return nil, err
	}

	deck := &game.Deck{}
	err = json.Unmarshal([]byte(jsonData), deck)
	if err != nil {
		return nil, err
	}

	return deck, nil
}

func (r RedisRepository) SaveDeck(deck *game.Deck) error {
	jsonData, err := json.Marshal(deck)
	if err != nil {
		return err
	}

	err = r.db.Set(r.ctx, deck.GetId().String(), jsonData, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
