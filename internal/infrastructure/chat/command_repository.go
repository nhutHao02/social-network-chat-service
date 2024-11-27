package chat

import (
	"github.com/nhutHao02/social-network-chat-service/config"
	"github.com/nhutHao02/social-network-chat-service/database"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/interface/chat"
)

type chatCommandRepository struct {
	cfg *config.Config
	db  *database.MongoDbClient
}

// ChatCommand implements chat.ChatCommandRepository.
func (c *chatCommandRepository) ChatCommand() {
	panic("unimplemented")
}

func NewChatCommandRepository(db *database.MongoDbClient, cfg *config.Config) chat.ChatCommandRepository {
	return &chatCommandRepository{db: db, cfg: cfg}
}
