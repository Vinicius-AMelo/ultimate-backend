package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
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

func InitQueue(queueName string) {
	waitForRabbitMQ()
	conn, err := amqp.Dial("amqp://guest:guest@rabbimq_service:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	q, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
}

func PublishMessage(queueName string, msg Message, replyQueue string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v", err)
	}

	err = ch.PublishWithContext(ctx, "", queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	log.Printf("Message sent: %s", body)
}

func WaitForResponseMessage(replyQueue string, chann chan string) (string, error) {
	msgs, err := ch.Consume(replyQueue, "", true, false, false, false, nil)
	if err != nil {
		return "fail", err
	}

	go func() {
		for msg := range msgs {

			var message Message
			if err := json.Unmarshal(msg.Body, &message); err != nil {
				chann <- "fail"
				return
			}
			log.Println("----", message)

			chann <- string(message.Value)
		}

	}()

	for {

		select {
		case msg := <-chann:
			log.Println("88888", msg)
			if msg == "fail" {
				return "fail", errors.New("failed to unmarshal message")
			}
			return msg, nil
		case <-time.After(5 * time.Second):
			return "fail", errors.New("timeout waiting for response")
		}
	}

}

func ConsumeMessages(queueName string, messagesChan chan Message) {
	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("er")
	}

	go func() {
		for d := range msgs {
			var message Message
			if err := json.Unmarshal(d.Body, &message); err != nil {
				log.Fatalf("Failed to unmarshal message: %v", err)
			}
			messagesChan <- message
		}
	}()

	select {}

}
