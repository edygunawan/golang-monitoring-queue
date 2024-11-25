package db

import (
	"context"
	"golang_api/app/helpers"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QueueMessage struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Message   string             `bson:"message"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func GetQueue() []QueueMessage {
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

	queueCollection := database.Collection("queueMessages")

	// Define a filter (empty to get all documents)
	filter := bson.D{{}}

	// Retrieve the documents from the collection
	cursor, err := queueCollection.Find(ctx, filter)
	if err != nil {
		log.Fatalf("Failed to find documents: %v", err)
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode each document
	var messages []QueueMessage
	for cursor.Next(ctx) {
		var msg QueueMessage
		if err := cursor.Decode(&msg); err != nil {
			log.Fatalf("Failed to decode document: %v", err)
		}
		messages = append(messages, msg)
	}

	// Check for any errors during iteration
	if err := cursor.Err(); err != nil {
		log.Fatalf("Cursor iteration error: %v", err)
	}

	return messages
}
