package chat

import (
	"context"

	"github.com/nhutHao02/social-network-chat-service/config"
	"github.com/nhutHao02/social-network-chat-service/database"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/entity"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/interface/chat"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type chatCommandRepository struct {
	cfg *config.Config
	db  *database.MongoDbClient
}

// InsertMessage implements chat.ChatCommandRepository.
func (repo *chatCommandRepository) InsertMessage(ctx context.Context, entityModel entity.Message) (string, error) {
	documentID, err := repo.db.InsertOne(ctx, repo.cfg.Database.DBName, entity.CollectionMessage, entityModel)
	if err != nil {
		logger.Error("chatCommandRepository-InsertMessage: Error inserting document", zap.Error(err))
		return "", err
	}
	return documentID.(primitive.ObjectID).Hex(), nil
}

func NewChatCommandRepository(db *database.MongoDbClient, cfg *config.Config) chat.ChatCommandRepository {
	return &chatCommandRepository{db: db, cfg: cfg}
}
