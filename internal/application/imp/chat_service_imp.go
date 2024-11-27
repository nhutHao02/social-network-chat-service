package imp

import (
	"github.com/nhutHao02/social-network-chat-service/internal/application"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/interface/chat"
	grpcUser "github.com/nhutHao02/social-network-user-service/pkg/grpc"
)

type chatService struct {
	queryRepo   chat.ChatQueryRepository
	commandRepo chat.ChatCommandRepository
	userClient  grpcUser.UserServiceClient
}

// Chat implements application.ChatService.
func (c *chatService) Chat() {
	panic("unimplemented")
}

func NewChatService(
	queryRepo chat.ChatQueryRepository,
	commandRepo chat.ChatCommandRepository,
	userClient grpcUser.UserServiceClient,
) application.ChatService {
	return &chatService{queryRepo: queryRepo, commandRepo: commandRepo, userClient: userClient}
}
