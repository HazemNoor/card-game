package player

func CreateNewPlayer(redisAddress, redisPassword string) *Player {
	return NewPlayer(NewRedisRepository(redisAddress, redisPassword))
}

func CreateNewPlayerInMemory() *Player {
	return NewPlayer(NewMemoryRepository())
}
