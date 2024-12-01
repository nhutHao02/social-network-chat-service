package chat

import (
	"context"

	"github.com/nhutHao02/social-network-chat-service/config"
	"github.com/nhutHao02/social-network-chat-service/database"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/entity"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/interface/chat"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type chatCommandRepository struct {
	cfg *config.Config
	db  *database.MongoDbClient
}

// UpdateRecentMessage implements chat.ChatCommandRepository.
func (repo *chatCommandRepository) UpdateRecentMessage(ctx context.Context, entityModel entity.Message) error {
	filter := bson.M{
		"$or": []bson.M{
			{
				"sender_id":   entityModel.SenderID,
				"receiver_id": entityModel.ReceiverID,
			},
			{
				"receiver_id": entityModel.SenderID,
				"sender_id":   entityModel.ReceiverID,
			},
		},
	}
	update := bson.M{
		"$set": entityModel,
	}
	opts := options.Update().SetUpsert(true)

	_, err := repo.db.UpdateOne(ctx, repo.cfg.Database.DBName, entity.CollectionRecentMessage, filter, update, opts)
	if err != nil {
		logger.Error("chatCommandRepository-InsertMessage: Error inserting document", zap.Error(err))
		return err
	}
	return nil
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
