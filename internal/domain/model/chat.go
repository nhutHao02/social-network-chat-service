package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetMessagesReq struct {
	SenderID   int64 `json:"senderID" form:"senderID"`
	ReceiverID int64 `json:"receiverID" form:"receiverID"`
	Page       int64 `form:"page"`
	Limit      int64 `form:"limit"`
	Token      string
}

type GetMessagesRes struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SenderID     int64              `bson:"sender_id" json:"senderID"`
	ReceiverID   int64              `bson:"receiver_id" json:"receiverID"`
	Content      string             `bson:"content" json:"content"`
	Timestamp    time.Time          `bson:"timestamp" json:"createdAt"`
	ReceiverInfo *UserInfo          `json:"info"`
}

type RecentMessageReq struct {
	UserID int64 `form:"userID"`
	Page   int64 `form:"page"`
	Limit  int64 `form:"limit"`
	Token  string
}

type RecentMessagesRes struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SenderID   int64              `bson:"sender_id" json:"senderID"`
	ReceiverID int64              `bson:"receiver_id" json:"receiverID"`
	Content    string             `bson:"content" json:"content"`
	Timestamp  time.Time          `bson:"timestamp" json:"createdAt"`
	Info       *UserInfo          `json:"info"`
}
