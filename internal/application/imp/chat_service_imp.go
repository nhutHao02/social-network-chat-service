package imp

import (
	"context"

	"github.com/nhutHao02/social-network-chat-service/internal/application"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/interface/chat"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/model"
	"github.com/nhutHao02/social-network-chat-service/pkg/constants"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	grpcUser "github.com/nhutHao02/social-network-user-service/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type chatService struct {
	queryRepo   chat.ChatQueryRepository
	commandRepo chat.ChatCommandRepository
	userClient  grpcUser.UserServiceClient
}

// GetMessages implements application.ChatService.
func (c *chatService) GetMessages(ctx context.Context, req model.GetMessagesReq) ([]model.GetMessagesRes, uint64, error) {
	res, total, err := c.queryRepo.GetMessages(ctx, req)
	if err != nil {
		return res, total, err
	}

	// Get User Info
	// Create context with metadata
	md := metadata.Pairs("authorization", constants.BearerString+req.Token)
	ctxx := metadata.NewOutgoingContext(ctx, md)

	userRes, err := c.userClient.GetUserInfo(ctxx, &grpcUser.GetUserRequest{UserID: req.ReceiverID})
	if err != nil {
		logger.Error("chatService-GetMessages: call grpcUser to server error", zap.Error(err))
		return res, total, err
	}

	for index, _ := range res {
		// pass user info to res
		res[index].ReceiverInfo = &model.UserInfo{
			ID:       int(userRes.Id),
			Email:    &userRes.Email,
			FullName: &userRes.FullName,
			UrlAvt:   &userRes.UrlAvt,
		}
	}

	return res, total, nil
}

func NewChatService(
	queryRepo chat.ChatQueryRepository,
	commandRepo chat.ChatCommandRepository,
	userClient grpcUser.UserServiceClient,
) application.ChatService {
	return &chatService{queryRepo: queryRepo, commandRepo: commandRepo, userClient: userClient}
}
