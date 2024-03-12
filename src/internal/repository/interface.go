package repository

import (
	"context"
	"grpc-test/internal/entity/user"
	"grpc-test/internal/transaction"
	"time"
)

type Info interface {
	SaveUser(ts transaction.Session, user user.User) error
}

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
}

type Logger interface{}
