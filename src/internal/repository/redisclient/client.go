package redisclient

import (
	"context"
	"grpc-test/internal/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(client *redis.Client) repository.RedisClient {
	return &redisClient{
		client: client,
	}
}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

