package chat

import (
	"github.com/nhutHao02/social-network-chat-service/config"
	"github.com/nhutHao02/social-network-chat-service/database"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/interface/chat"
)

type chatQueryRepository struct {
	cfg *config.Config
	db  *database.MongoDbClient
}

// ChatQuery implements chat.ChatQueryRepository.
func (c *chatQueryRepository) ChatQuery() {
	panic("unimplemented")
}

func NewChatQueryRepository(db *database.MongoDbClient, cfg *config.Config) chat.ChatQueryRepository {
	return &chatQueryRepository{db: db, cfg: cfg}
}
