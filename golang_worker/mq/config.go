package mq

import (
	"log"

	"golang_worker/app/helpers"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Connect to RabbitMQ
func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	rabbitMQString := helpers.GoDotEnvVariable("AMQP_URL")
	// Establish connection to RabbitMQ
	connection, err := amqp.Dial(rabbitMQString)
	if err != nil {
		return nil, nil, err // Return error if connection fails
	}

	// Create a channel
	channel, err := connection.Channel()
	if err != nil {
		connection.Close() // Close connection if channel creation fails
		return nil, nil, err
	}

	return connection, channel, nil // Return connection and channel along with nil error
}

// Close the connection
func CloseRabbitMQ(connection *amqp.Connection, channel *amqp.Channel) {
	if err := channel.Close(); err != nil {
		log.Fatal("Failed to close channel:", err)
	}
	if err := connection.Close(); err != nil {
		log.Fatal("Failed to close connection:", err)
	}
}
