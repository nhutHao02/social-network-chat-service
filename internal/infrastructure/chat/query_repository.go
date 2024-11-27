package chat

import (
	"context"

	"github.com/nhutHao02/social-network-chat-service/config"
	"github.com/nhutHao02/social-network-chat-service/database"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/entity"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/interface/chat"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/model"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type chatQueryRepository struct {
	cfg *config.Config
	db  *database.MongoDbClient
}

// GetMessages implements chat.ChatQueryRepository.
func (repo *chatQueryRepository) GetMessages(ctx context.Context, req model.GetMessagesReq) ([]model.GetMessagesRes, uint64, error) {
	var res []model.GetMessagesRes

	filter := bson.M{
		"$or": []bson.M{
			{
				"sender_id":   req.SenderID,
				"receiver_id": req.ReceiverID,
			},
			{
				"receiver_id": req.SenderID,
				"sender_id":   req.ReceiverID,
			},
		},
	}

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "timestamp", Value: -1}})
	opts.SetSkip((req.Page - 1) * req.Limit)
	opts.SetLimit(req.Limit)

	err := repo.db.FindMany(ctx, repo.cfg.Database.DBName, entity.CollectionMessage, filter, &res, opts)
	if err != nil {
		logger.Error("chatQueryRepository-GetMessages: FindMany message error", zap.Error(err))
		return res, 0, err
	}

	totalCount, err := repo.db.CountDocuments(ctx, repo.cfg.Database.DBName, entity.CollectionMessage, filter)
	if err != nil {
		logger.Error("chatQueryRepository-GetMessages: Count Documents error", zap.Error(err))
		return res, 0, err
	}

	return res, uint64(totalCount), nil
}

func NewChatQueryRepository(db *database.MongoDbClient, cfg *config.Config) chat.ChatQueryRepository {
	return &chatQueryRepository{db: db, cfg: cfg}
}
