package mq

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Send a message to the RabbitMQ queue
func ConsumeMessages(channel *amqp.Channel, queueName string, autoAck bool) ([]MessageBody, error) {
	// Declare the queue to make sure it exists
	_, err := channel.QueueDeclare(
		queueName, // queue name
		true,      // durable (the queue will survive server restarts)
		false,     // delete when unused
		false,     // exclusive (only this connection can access the queue)
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatal("Error declaring queue:", err)
	}

	// Start consuming messages from the queue
	msgs, err := channel.Consume(
		queueName, // queue
		"",        // consumer
		autoAck,   // auto-ack (set to false to manually acknowledge messages)
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Failed to register a consumer: %v", err)
	}

	var messages []MessageBody
	var wg sync.WaitGroup
	for msg := range msgs {
		wg.Add(1)
		go func(msg amqp.Delivery) {
			defer wg.Done()
			var data MessageBody
			if err := json.Unmarshal([]byte(msg.Body), &data); err != nil {
				log.Printf("Failed to parse queue message: %v", msg.Body)
			} else {
				log.Printf("data: %v", data)
				messages = append(messages, data)
			}
		}(msg)
	}
	wg.Wait()

	log.Printf("messages: %v", messages)
	return messages, nil
}
