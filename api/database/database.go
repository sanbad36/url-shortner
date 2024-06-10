// database/database.go

package database

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sanbad36/url-shortner/api/database/store"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

func Init() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Println("Failed to connect to Redis, using in-memory store:", err)
		redisClient = nil
	}
}

func Set(key, value string, expiration time.Duration) error {
	if redisClient != nil {
		return redisClient.Set(ctx, key, value, expiration).Err()
	}

	store.Set(key, value, expiration)
	return nil
}

func Get(key string) (string, error) {
	if redisClient != nil {
		return redisClient.Get(ctx, key).Result()
	}

	val, exists := store.Get(key)
	if !exists {
		return "", errors.New("key not found")
	}
	return val, nil
}

func Delete(key string) error {
	if redisClient != nil {
		return redisClient.Del(ctx, key).Err()
	}

	store.Delete(key)
	return nil
}

func Keys(pattern string) ([]string, error) {
	if redisClient != nil {
		return redisClient.Keys(ctx, pattern).Result()
	}

	return store.Keys(pattern)
}

func Decr(key string) error {
	if redisClient != nil {
		return redisClient.Decr(ctx, key).Err()
	}

	return store.Decr(key)
}

func TTL(key string) (time.Duration, error) {
	if redisClient != nil {
		return redisClient.TTL(ctx, key).Result()
	}

	return store.TTL(key)
}

func IsRedisAvailable() bool {
	return redisClient != nil
}
