package application

import (
	"context"

	"github.com/nhutHao02/social-network-chat-service/internal/domain/model"
)

type ChatService interface {
	GetMessages(ctx context.Context, req model.GetMessagesReq) ([]model.GetMessagesRes, uint64, error)
}
