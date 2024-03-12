package rimport

import "grpc-test/internal/repository"

type Repository struct {
	Logger      repository.Logger
	Info        repository.Info
	RedisClient repository.RedisClient
}

type MockRepository struct {
	Info        *repository.MockInfo
	Logger      *repository.MockLogger
	RedisClient *repository.MockRedisClient
}
