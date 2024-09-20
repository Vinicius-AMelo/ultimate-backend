package main

import (
	"context"
	"fmt"
	"time"

	"github.com/my/repo/services/api/rabbitmq"
	"github.com/redis/go-redis/v9"
)

func main() {
	rabbitmq.InitQueue()
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := client.Set(ctx, "key", "value", 10*time.Minute).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
