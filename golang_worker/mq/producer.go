package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// Send a message to the RabbitMQ queue
func ProduceMessage(channel *amqp.Channel, queueName string, message []byte) error {
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
		return err
	}

	// Send the message
	err = channel.Publish(
		"",        // Exchange
		queueName, // Routing key (queue name)
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	return err
}
