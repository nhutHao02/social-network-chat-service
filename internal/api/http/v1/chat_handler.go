package v1

import (
	"github.com/nhutHao02/social-network-chat-service/internal/application"
)

type ChatHandler struct {
	chatService application.ChatService
}

func NewChatHandler(chatService application.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}
