package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

var ch *amqp.Channel
var q amqp.Queue

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

func InitQueue() {
	waitForRabbitMQ()
	conn, err := amqp.Dial("amqp://guest:guest@rabbimq_service:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	q, err = ch.QueueDeclare("hello", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
}

func PublishMessage(ctx context.Context, msg Message) {
	body, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v", err)
	}

	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	log.Printf("Message sent: %s", body)
}
