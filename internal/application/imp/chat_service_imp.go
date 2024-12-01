package imp

import (
	"context"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/nhutHao02/social-network-chat-service/internal/application"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/entity"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/interface/chat"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/model"
	"github.com/nhutHao02/social-network-chat-service/pkg/constants"
	ws "github.com/nhutHao02/social-network-chat-service/pkg/websocket"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	grpcUser "github.com/nhutHao02/social-network-user-service/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type chatService struct {
	queryRepo   chat.ChatQueryRepository
	commandRepo chat.ChatCommandRepository
	userClient  grpcUser.UserServiceClient
	ws          *ws.Socket
}

// PrivateMessageWS implements application.ChatService.
func (c *chatService) PrivateMessageWS(ctx context.Context, conn *websocket.Conn, req model.MessageReq) {
	// create roomID, userID connect and model
	var messageRoomID string
	if req.SenderID > req.ReceiverID {
		messageRoomID = strconv.Itoa(int(req.SenderID)) + strconv.Itoa(int(req.ReceiverID)) + "message"
	} else {
		messageRoomID = strconv.Itoa(int(req.ReceiverID)) + strconv.Itoa(int(req.SenderID)) + "message"
	}
	var (
		userWSID        = strconv.Itoa(int(req.SenderID))
		incomingMessage = model.IncomingMessageWSReq{}
	)

	// Add connection
	c.ws.AddConnection(messageRoomID, userWSID, conn)

	// Listen for connection close event
	defer c.ws.RemoveConnection(messageRoomID, userWSID, conn)

	// Get User Info
	// Create context with metadata
	md := metadata.Pairs("authorization", constants.BearerString+req.Token)
	ctxx := metadata.NewOutgoingContext(ctx, md)

	userRes, err := c.userClient.GetUserInfo(ctxx, &grpcUser.GetUserRequest{UserID: req.SenderID})
	if err != nil {
		logger.Error("tweetService-CommentWebSocket: Error get UserInfo, call grpcUser to server error", zap.Error(err))
	}

	// Handle incoming messages
	for {
		if err := conn.ReadJSON(&incomingMessage); err != nil {
			logger.Error("chatService-PrivateMessageWS: Error reading message", zap.Error(err))
			c.ws.RemoveConnection(messageRoomID, userWSID, conn)
			break
		}
		currentTime := time.Now()
		// Save message to db
		entityModel := entity.Message{
			SenderID:   req.SenderID,
			ReceiverID: req.ReceiverID,
			Content:    incomingMessage.Message,
			Timestamp:  currentTime,
		}

		documentID, err := c.commandRepo.InsertMessage(ctx, entityModel)
		if err != nil {
			logger.Error("tweetService-CommentWebSocket: Error saving comment to DB and ignore broadcast", zap.Error(err))
			continue
		}

		if err = c.commandRepo.UpdateRecentMessage(ctx, entityModel); err != nil {
			logger.Error("tweetService-CommentWebSocket: Error update recent message to DB and ignore broadcast", zap.Error(err))
			continue
		}

		outGoingMessage := model.OutgoingMessageWSRes{
			ID: documentID,
			Sender: &model.UserInfo{
				ID:       int(userRes.Id),
				Email:    &userRes.Email,
				FullName: &userRes.FullName,
				UrlAvt:   &userRes.UrlAvt,
			},
			Message:   incomingMessage.Message,
			CreatedAt: currentTime,
		}
		// Broadcast message to the room
		c.ws.Broadcast(messageRoomID, userWSID, outGoingMessage)
	}

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
	ws *ws.Socket,
) application.ChatService {
	return &chatService{queryRepo: queryRepo, commandRepo: commandRepo, userClient: userClient, ws: ws}
}
