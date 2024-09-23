package main

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	rabbitmq "github.com/my/repo/services/rabbitmq"
	"github.com/redis/go-redis/v9"
)

type session struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

var client *redis.Client

func main() {
	messagesChan := make(chan rabbitmq.Message)
	rabbitmq.InitQueue("redis")
	// rabbitmq.InitQueue("respondToApi")
	var wg sync.WaitGroup

	client = initClient()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for msg := range messagesChan {
			if err := processMessages(msg); err != nil {
				log.Println("Failed to process message")
			}

		}
	}()

	rabbitmq.ConsumeMessages("redis", messagesChan)

	wg.Wait()
}

func processMessages(msg rabbitmq.Message) error {
	var body session

	if err := json.Unmarshal(msg.Value, &body); err != nil {
		return err
	}

	log.Println(msg.Key)

	if msg.Key == "setToken" {
		if err := setToken(context.Background(), body); err != nil {
			rabbitmq.PublishMessage("responseToApi", rabbitmq.Message{Key: "fail", Value: json.RawMessage(err.Error())})
		}

		rabbitmq.PublishMessage("responseToApi", rabbitmq.Message{Key: "success", Value: json.RawMessage("")})
	} else if msg.Key == "getToken" {
		if token := getToken(context.Background(), body); token == nil {
			rabbitmq.PublishMessage("responseToApi", rabbitmq.Message{Key: "fail", Value: json.RawMessage("invalid token")})
		} else {
			rabbitmq.PublishMessage("responseToApi", rabbitmq.Message{Key: "success", Value: json.RawMessage(token.Val())})
		}

	}
	return nil
}

func initClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func setToken(ctx context.Context, session session) error {
	if err := client.Set(ctx, session.ID, session.Token, 10*time.Minute).Err(); err != nil {
		return err
	}

	return nil
}

func getToken(ctx context.Context, session session) *redis.StringCmd {
	return client.Get(ctx, session.ID)

}
