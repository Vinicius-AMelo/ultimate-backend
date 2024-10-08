package redisClient

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func InitClient() {

	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

}

func SetToken(ctx context.Context, id string, token string) error {
	if err := client.Set(ctx, id, token, 10*time.Minute).Err(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func GetToken(ctx context.Context, id string) string {
	return client.Get(ctx, id).Val()

}
