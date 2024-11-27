package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nhutHao02/social-network-chat-service/config"
	v1 "github.com/nhutHao02/social-network-chat-service/internal/api/http/v1"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"go.uber.org/zap"
)

type HTTPServer struct {
	Cfg *config.Config
	// handlers
	ChatHandler *v1.ChatHandler
}

func NewHTTPServer(cfg *config.Config, tweetHandler *v1.ChatHandler) *HTTPServer {
	return &HTTPServer{Cfg: cfg, ChatHandler: tweetHandler}
}

func (s *HTTPServer) RunHTTPServer() error {
	r := gin.Default()
	v1.MapRoutes(r, s.ChatHandler)
	logger.Info("HTTP Server server listening at" + s.Cfg.HTTPServer.Address)
	err := r.Run(s.Cfg.HTTPServer.Address)
	if err != nil {
		logger.Error("HTTP Server error", zap.Error(err))
		return err
	}
	return nil
}
