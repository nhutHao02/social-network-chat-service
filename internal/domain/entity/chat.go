package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	SenderID   int64              `bson:"sender_id"`
	ReceiverID int64              `bson:"receiver_id"`
	Content    string             `bson:"content"`
	Timestamp  time.Time          `bson:"timestamp"`
}
