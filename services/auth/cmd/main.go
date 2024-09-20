package main

import (
	"encoding/json"
	"log"
	"sync"

	rabbitmq "github.com/my/repo/services/rabbitmq"
)

type session struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func main() {
	messagesChan := make(chan rabbitmq.Message)
	rabbitmq.InitQueue("redis")
	// rabbitmq.InitQueue("respondToApi")
	var wg sync.WaitGroup

	// ctx := context.Background()

	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "redis:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	wg.Add(1)
	go func() {
		defer wg.Done()
		var body session
		for msg := range messagesChan {
			if err := json.Unmarshal(msg.Value, &body); err != nil {
				log.Println("Failed to unmarshall json")
			}

		}
	}()

	rabbitmq.ConsumeMessages("redis", messagesChan)

	wg.Wait()
}
