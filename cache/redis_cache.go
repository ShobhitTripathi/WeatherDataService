package cache

import (
	"WeatherDataService/model"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	fmt.Printf("Redis Cache initialized")
	return &RedisCache{client: rdb}
}

func (rc *RedisCache) Get(ctx context.Context, key string) (*model.CacheEntry, error) {
	val, err := rc.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var entry model.CacheEntry
	err = json.Unmarshal([]byte(val), &entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (rc *RedisCache) Set(ctx context.Context, key string, entry *model.CacheEntry, expiration time.Duration) error {
	val, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	err = rc.client.Set(ctx, key, val, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}
