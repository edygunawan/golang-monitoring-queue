package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_worker/app/helpers"
	"golang_worker/app/mq"
	"golang_worker/app/queue"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	/* Mongodb */
	dbConnectionString := helpers.GoDotEnvVariable("DB_CONN_STRING")
	databaseName := helpers.GoDotEnvVariable("DB_NAME")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbConnectionString))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	// Ping the database to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Access the database
	database := client.Database(databaseName)
	log.Println("Connected to database:", databaseName)

	/* RabbitMQ */
	connection, channel, errMQ := mq.ConnectRabbitMQ()
	if errMQ != nil {
		log.Fatal(errMQ)
	}
	defer mq.CloseRabbitMQ(connection, channel)
	queueName := helpers.GoDotEnvVariable("MQ_QUEUE")
	msgs := mq.ConsumeMessages(channel, queueName, true)

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			/* Ready data from queue */
			var data queue.QueueMessage
			if err := json.Unmarshal([]byte(msg.Body), &data); err != nil {
				log.Printf("Failed to parse queue message: %v", msg.Body)
			} else {
				log.Printf("messages: %v", data)
				/* Insert to database */
				nowTime := time.Now()
				data.CreatedAt = nowTime
				data.UpdatedAt = nowTime

				queueCollection := database.Collection("queueMessages")
				insertResult, err := queueCollection.InsertOne(context.TODO(), data)
				if err != nil {
					log.Printf("Failed to insert queue message: %v", err)
					/* Send it bakc to queue if insert failed */
					errPush := mq.ProduceMessage(channel, queueName, msg.Body)
					if errPush != nil {
						log.Println("Failed to re-publish message:", errPush)
					}
				} else {
					fmt.Println("Queue executed", insertResult.InsertedID)
				}
			}
		}
	}()

	log.Printf(" [*] Waiting for queue. To exit press CTRL+C")
	<-forever
}
