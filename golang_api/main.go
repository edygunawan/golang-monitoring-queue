package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "golang_api/app/mq" // Import the rabbitmq package
    "log"
)

func main() {
    // Connect to RabbitMQ
    rabbitmq.ConnectRabbitMQ()
    defer rabbitmq.CloseRabbitMQ() // Ensure the connection is closed when the app exits

    router := gin.Default()

    // Define a route for GET request
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "ok!",
        })
    })

    // Define a route for POST request to send a message to RabbitMQ
    router.POST("/queue", func(c *gin.Context) {
        // Here you would get the message from the request body or query params
        var json struct {
            Message string `json:"message"`
        }

        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid request",
            })
            return
        }

        // Send the message to RabbitMQ
        err := rabbitmq.ProduceMessage("my_queue", json.Message)
        if err != nil {
            log.Println("Failed to publish message:", err)
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to send message to queue",
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Message sent to queue successfully",
        })
    })

    // Start the server on port 8080
    router.Run(":8080")
}
