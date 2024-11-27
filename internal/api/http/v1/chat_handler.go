package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nhutHao02/social-network-chat-service/internal/application"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/model"
	"github.com/nhutHao02/social-network-chat-service/pkg/constants"
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
//	@Router			/guest/sign-up [post]
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
