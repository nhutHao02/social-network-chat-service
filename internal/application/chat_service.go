package application

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/model"
)

type ChatService interface {
	GetMessages(ctx context.Context, req model.GetMessagesReq) ([]model.GetMessagesRes, uint64, error)
	PrivateMessageWS(ctx context.Context, conn *websocket.Conn, req model.MessageReq)
}
