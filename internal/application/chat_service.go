package application

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/model"
)

type ChatService interface {
	GetMessages(ctx context.Context, req model.GetMessagesReq) ([]model.GetMessagesRes, uint64, error)
	PrivateMessageWS(ctx context.Context, conn *websocket.Conn, req model.MessageReq)
	GetRecentMessage(ctx context.Context, req model.RecentMessageReq) ([]model.RecentMessagesRes, uint64, error)
	RecentMessageWS(ctx context.Context, conn *websocket.Conn, req model.WSRecentReq)
}
