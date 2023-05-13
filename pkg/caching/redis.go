package caching

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type ICache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (*string, error)
}

type cache struct {
	client *redis.Client
}

func NewCache() ICache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &cache{client: client}
}

func (c *cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := c.client.Set(ctx, key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func (c *cache) Get(ctx context.Context, key string) (*string, error) {
	value, err := c.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &value, nil
}
