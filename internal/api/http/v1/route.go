package v1

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/nhutHao02/social-network-chat-service/docs"
	"github.com/nhutHao02/social-network-common-service/middleware"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapRoutes(
	router *gin.Engine,
	chatHandler *ChatHandler,
) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ws/messages", chatHandler.MessageWSHandler)
		v1.GET("/ws/recent", chatHandler.MessageWSRecentHandler)
		v1.Use(middleware.JwtAuthMiddleware(logger.GetDefaultLogger()))
		{
			vChat := v1.Group("/chat")
			vChat.GET("", chatHandler.GetPrivateMessages)
			vChat.GET("/recent", chatHandler.GetRecentMessage)

			vSocket := v1.Group("/ws")
			vSocket.GET("private-message", chatHandler.MessageWebSocketHandler)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
