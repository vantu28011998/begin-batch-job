package redis

import (
	"os"

	"github.com/go-redis/redis"
)

func GetRedisConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return client
}
