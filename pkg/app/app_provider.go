package app

func CreateNewApp(redisAddress, redisPassword string) *App {
	return NewApp(NewRedisRepository(redisAddress, redisPassword))
}

func CreateNewAppInMemory() *App {
	return NewApp(NewMemoryRepository())
}
