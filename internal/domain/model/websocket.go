package model

import "time"

type MessageReq struct {
	SenderID   int64  `form:"senderID"`
	ReceiverID int64  `form:"receiverID"`
	Token      string `form:"token"`
}

type IncomingMessageWSReq struct {
	Message string `json:"message"`
}

type OutgoingMessageWSRes struct {
	ID         string    `json:"id"`
	Sender     *UserInfo `json:"info"`
	Message    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	SenderID   int64     `json:"senderID"`
	ReceiverID int64     `json:"receiverID"`
}
