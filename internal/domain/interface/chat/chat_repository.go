package chat

import (
	"context"

	"github.com/nhutHao02/social-network-chat-service/internal/domain/entity"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/model"
)

type ChatQueryRepository interface {
	GetMessages(ctx context.Context, req model.GetMessagesReq) ([]model.GetMessagesRes, uint64, error)
}

type ChatCommandRepository interface {
	InsertMessage(ctx context.Context, entityModel entity.Message) (string, error)
}
