package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type Message struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

type session struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func waitForRabbitMQ() {
	for {
		conn, err := amqp.Dial("amqp://guest:guest@rabbimq_service:5672/")
		if err == nil {
			conn.Close()
			break
		}
		log.Println("Waiting for RabbitMQ to be available...")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	waitForRabbitMQ()
	ctx := context.Background()

	conn, err := amqp.Dial("amqp://guest:guest@rabbimq_service:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	go func() {
		for d := range msgs {
			var message Message
			if err := json.Unmarshal(d.Body, &message); err != nil {
				panic(err)
			}

			if message.Key == "redis" {
				var body session
				if err := json.Unmarshal([]byte(message.Value), &body); err != nil {
					panic(err)
				}

				if err := client.Set(ctx, body.ID, body.Token, 10*time.Minute).Err(); err != nil {
					panic(err)
				}
				log.Printf("Session stored in Redis: ID=%s, Token=%s", body.ID, body.Token)
			}
		}
	}()

	select {}
}
