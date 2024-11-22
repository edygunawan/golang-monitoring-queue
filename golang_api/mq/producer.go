package mq

import (
    "log"
    "github.com/rabbitmq/amqp091-go"
)

var connection *amqp091-go.Connection
var channel *amqp091-go.Channel

// Connect to RabbitMQ
func ConnectRabbitMQ() {
    var err error
    // Update the connection string if needed (username/password/hostname)
    connection, err = amqp091-go.Dial("amqp://user:password@rabbitmq:5672/")
    if err != nil {
        log.Fatal("Failed to connect to RabbitMQ:", err)
    }

    // Create a channel
    channel, err = connection.Channel()
    if err != nil {
        log.Fatal("Failed to open a channel:", err)
    }
}

// Close the connection
func CloseRabbitMQ() {
    if err := channel.Close(); err != nil {
        log.Fatal("Failed to close channel:", err)
    }
    if err := connection.Close(); err != nil {
        log.Fatal("Failed to close connection:", err)
    }
}

// Send a message to the RabbitMQ queue
func ProduceMessage(queueName string, message string) error {
    // Declare the queue
    _, err := channel.QueueDeclare(
        queueName, // Queue name
        false,     // Durable (doesn't survive broker restart)
        false,     // Delete when unused
        false,     // Exclusive (only this connection can access)
        false,     // No-wait
        nil,       // Arguments
    )
    if err != nil {
        return err
    }

    // Send the message
    err = channel.Publish(
        "",         // Exchange
        queueName,  // Routing key (queue name)
        false,      // Mandatory
        false,      // Immediate
        amqp091-go.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        },
    )
    return err
}
