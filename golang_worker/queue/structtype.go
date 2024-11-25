package queue

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QueueMessage struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Message   string             `bson:"message"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
