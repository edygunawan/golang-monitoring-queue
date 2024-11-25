package mq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Send a message to the RabbitMQ queue
func ConsumeMessages(channel *amqp.Channel, queueName string, autoAck bool) <-chan amqp.Delivery {

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
	}

	return msgs
}
