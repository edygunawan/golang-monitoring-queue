package main

import (
	"golang_api/app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define a route for GET request
	router.GET("/", controllers.Status)
	router.POST("/queue", controllers.PostQueueMessage)
	router.GET("/queue", controllers.GetQueues)
	router.GET("/queue-messages", controllers.GetQueueMessages)

	// Start the server on port 8080
	router.Run(":8080")
}
