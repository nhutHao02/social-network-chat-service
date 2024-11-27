package model

import "time"

type MessageReq struct {
	SenderID   int64 `form:"senderID"`
	ReceiverID int64 `form:"receiverID"`
	Token      string
}

type IncomingMessageWSReq struct {
	Message string `json:"message"`
}

type OutgoingMessageWSRes struct {
	ID        string    `json:"id"`
	Sender    *UserInfo `json:"sender"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}
