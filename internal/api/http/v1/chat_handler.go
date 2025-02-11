package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nhutHao02/social-network-chat-service/internal/application"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/model"
	"github.com/nhutHao02/social-network-chat-service/pkg/constants"
	"github.com/nhutHao02/social-network-chat-service/pkg/websocket"
	common "github.com/nhutHao02/social-network-common-service/model"
	"github.com/nhutHao02/social-network-common-service/request"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-common-service/utils/token"
	"go.uber.org/zap"
)

type ChatHandler struct {
	chatService application.ChatService
}

func NewChatHandler(chatService application.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}

// GetPrivateMessages godoc
//
//	@Summary		GetPrivateMessages
//	@Description	Get Messages between two users
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			senderID	query		int												true	"SenderID"
//	@Param			receiverID	query		int												true	"ReceiverID"
//	@Param			page		query		int												false	"Page"
//	@Param			limit		query		int												false	"Limit"
//	@Success		200			{object}	common.Response{data=[]model.GetMessagesRes}	"successfully"
//	@Failure		default		{object}	common.Response{data=nil}						"failure"
//	@Router			/chat [get]
func (h *ChatHandler) GetPrivateMessages(c *gin.Context) {
	var req model.GetMessagesReq

	err := request.GetQueryParamsFromUrl(c, &req)
	if err != nil {
		return
	}
	userID, err := token.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.GetPrivateMessagesFailure))
		return
	}

	if userID != int(req.SenderID) {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(constants.InvalidUserID, constants.GetPrivateMessagesFailure))
		return
	}

	token, err := token.GetTokenString(c)
	if err != nil {
		logger.Error("ChatHandler-GetPrivateMessages: get token from request error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error(), constants.GetPrivateMessagesFailure))
		return
	}

	req.Token = token
	res, total, err := h.chatService.GetMessages(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.GetPrivateMessagesFailure))
		return
	}

	c.JSON(http.StatusOK, common.NewPagingSuccessResponse(res, total))
}

// MessageWebSocketHandler godoc
//
//	@Summary		MessageWebSocketHandler
//	@Description	Establish a WebSocket connection to send messages between users in real-time.
//	@Tags			Tweet
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer <your_token>"
//	@Param			userID			query		int							true	"User ID"
//	@Param			tweetID			query		int							true	"Tweet ID"
//	@Success		101				{string}	string						"WebSocket connection established"
//	@Failure		default			{object}	common.Response{data=nil}	"failure"
//	@Router			/ws/private-message [get]
func (h *ChatHandler) MessageWebSocketHandler(c *gin.Context) {
	var req model.MessageReq

	if err := request.GetQueryParamsFromUrl(c, &req); err != nil {
		return
	}

	userID, err := token.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.ConnectPrivateMessageWebSocketFailure))
		return
	}

	if userID != int(req.SenderID) {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(constants.InvalidUserID, constants.ConnectPrivateMessageWebSocketFailure))
		return
	}

	token, err := token.GetTokenString(c)
	if err != nil {
		logger.Error("ChatHandler-MessageWebSocketHandler: get token from request error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error(), constants.ConnectPrivateMessageWebSocketFailure))
		return
	}

	req.Token = token

	// Upgrade HTTP connection to WebSocket
	conn, err := websocket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("Error when upgrade HTTP connection to WebSocket", zap.Error(err))
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error(), constants.ConnectPrivateMessageWebSocketFailure))
		return
	}

	h.chatService.PrivateMessageWS(c.Request.Context(), conn, req)
}

// MessageWSHandler godoc
//
//	@Summary		MessageWSHandler
//	@Description	Establish a WebSocket connection to send messages between users in real-time.
//	@Tags			Tweet
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer <your_token>"
//	@Param			userID			query		int							true	"User ID"
//	@Param			tweetID			query		int							true	"Tweet ID"
//	@Success		101				{string}	string						"WebSocket connection established"
//	@Failure		default			{object}	common.Response{data=nil}	"failure"
//	@Router			/ws/private-message [get]
func (h *ChatHandler) MessageWSHandler(c *gin.Context) {
	var req model.MessageReq

	if err := request.GetQueryParamsFromUrl(c, &req); err != nil {
		return
	}

	// Upgrade HTTP connection to WebSocket
	conn, err := websocket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("Error when upgrade HTTP connection to WebSocket", zap.Error(err))
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error(), constants.ConnectPrivateMessageWebSocketFailure))
		return
	}

	h.chatService.PrivateMessageWS(c.Request.Context(), conn, req)
}

// MessageWSRecentHandler godoc
//
//	@Summary		MessageWSRecentHandler
//	@Description	Establish a WebSocket connection to send recent messages to user in real-time.
//	@Tags			Tweet
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer <your_token>"
//	@Param			userID			query		int							true	"User ID"
//	@Param			tweetID			query		int							true	"Tweet ID"
//	@Success		101				{string}	string						"WebSocket connection established"
//	@Failure		default			{object}	common.Response{data=nil}	"failure"
//	@Router			/ws/private-message [get]
func (h *ChatHandler) MessageWSRecentHandler(c *gin.Context) {
	var req model.WSRecentReq

	if err := request.GetQueryParamsFromUrl(c, &req); err != nil {
		return
	}

	// Upgrade HTTP connection to WebSocket
	conn, err := websocket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("Error when upgrade HTTP connection to WebSocket", zap.Error(err))
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error(), constants.ConnectPrivateMessageWebSocketFailure))
		return
	}

	h.chatService.RecentMessageWS(c.Request.Context(), conn, req)
}

// GetRecentMessage godoc
//
//	@Summary		GetRecentMessage
//	@Description	Get Recent Message by userID
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			userID	query		int												true	"UserID"
//	@Param			page	query		int												false	"Page"
//	@Param			limit	query		int												false	"Limit"
//	@Success		200		{object}	common.Response{data=[]model.GetMessagesRes}	"successfully"
//	@Failure		default	{object}	common.Response{data=nil}						"failure"
//	@Router			/chat/recent [get]
func (h *ChatHandler) GetRecentMessage(c *gin.Context) {
	var req model.RecentMessageReq

	if err := request.GetQueryParamsFromUrl(c, &req); err != nil {
		return
	}

	userID, err := token.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.GetRecentMessagesFailure))
		return
	}

	if userID != int(req.UserID) {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(constants.InvalidUserID, constants.GetRecentMessagesFailure))
		return
	}

	token, err := token.GetTokenString(c)
	if err != nil {
		logger.Error("ChatHandler-GetRecentMessage: get token from request error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error(), constants.GetRecentMessagesFailure))
		return
	}

	req.Token = token

	res, total, err := h.chatService.GetRecentMessage(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.GetRecentMessagesFailure))
		return
	}

	c.JSON(http.StatusOK, common.NewPagingSuccessResponse(res, total))
}
