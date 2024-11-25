package controllers

import (
	"encoding/json"
	"golang_api/app/db"
	"golang_api/app/mq"
	"log"
	"net/http"

	"golang_api/app/helpers"

	"github.com/gin-gonic/gin"
)

type QueueDto struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// GetRabbitMQQueues fetches the queue list from RabbitMQ
func GetQueues(c *gin.Context) {
	rabbitMQAPI := helpers.GoDotEnvVariable("MQ_MANAGEMENT_URL")
	username := helpers.GoDotEnvVariable("MQ_USERNAME")
	password := helpers.GoDotEnvVariable("MQ_PASSWORD")

	// Create an HTTP client
	req, err := http.NewRequest("GET", rabbitMQAPI, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.SetBasicAuth(username, password)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch queues"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch queues from RabbitMQ"})
		return
	}

	// Parse the response
	var queues []mq.RabbitMQQueue
	if err := json.NewDecoder(resp.Body).Decode(&queues); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	// Return the list of queues
	c.JSON(http.StatusOK, gin.H{"queues": queues})
}

func PostQueueMessage(c *gin.Context) {
	connection, channel, errMQ := mq.ConnectRabbitMQ()
	if errMQ != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Server error",
		})
		return
	}

	defer mq.CloseRabbitMQ(connection, channel)

	var message mq.MessageBody
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	jsonData, errM := json.Marshal(message)
	if errM != nil {
		log.Fatalf("Failed to marshal message: %v", errM)
	}

	log.Printf("json: %v", message)
	log.Printf("json mq.MessageBody: %v", jsonData)

	// Send the message to RabbitMQ
	queueName := helpers.GoDotEnvVariable("MQ_QUEUE")
	err := mq.ProduceMessage(channel, queueName, jsonData)
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
}

func GetQueueMessages(c *gin.Context) {
	/* Get saved queue from database */
	messages := db.GetQueue()
	log.Printf("Queue messages: %v", messages)

	var queueMessages []QueueDto
	for _, message := range messages {
		data := QueueDto{
			Message: message.Message,
			Status:  "done",
		}
		queueMessages = append(queueMessages, data)
	}

	/* Example failed message that will pushed back to queue  from nextjs app */
	failedQueue := QueueDto{
		Message: "Nest message",
		Status:  "failed",
	}
	queueMessages = append(queueMessages, failedQueue)

	// Return the list of messages as a JSON response
	c.JSON(http.StatusOK, queueMessages)
}
