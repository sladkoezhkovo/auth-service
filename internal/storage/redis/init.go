package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sladkoezhkovo/auth-service/internal/configs"
	"golang.org/x/net/context"
	"os"
	"time"
)

type redisStorage struct {
	client *redis.Client
}

func New(config *configs.RedisConfig) *redisStorage {
	return &redisStorage{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       config.Db,
		}),
	}
}

func (r *redisStorage) Set(key string, value interface{}, ttl int) error {
	ctx := context.Background()
	return r.client.Set(ctx, key, value, time.Duration(ttl)).Err()
}

func (r *redisStorage) Get(key string) (string, error) {
	ctx := context.Background()
	return r.client.Get(ctx, key).Result()
}

func (r *redisStorage) Clear(key string) error {
	ctx := context.Background()
	return r.client.Del(ctx, key).Err()
}
