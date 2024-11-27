//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/nhutHao02/social-network-chat-service/config"
	"github.com/nhutHao02/social-network-chat-service/database"
	"github.com/nhutHao02/social-network-chat-service/internal/api"
	"github.com/nhutHao02/social-network-chat-service/internal/api/http"
	"github.com/nhutHao02/social-network-chat-service/internal/api/http/v1"
	"github.com/nhutHao02/social-network-chat-service/internal/application/imp"
	"github.com/nhutHao02/social-network-chat-service/internal/infrastructure/chat"
	"github.com/nhutHao02/social-network-chat-service/pkg/redis"
	ws "github.com/nhutHao02/social-network-chat-service/pkg/websocket"
	grpcUser "github.com/nhutHao02/social-network-user-service/pkg/grpc"
)

var serverSet = wire.NewSet(
	api.NewSerVer,
)

var itemServerSet = wire.NewSet(
	http.NewHTTPServer,
)

var httpHandlerSet = wire.NewSet(
	v1.NewChatHandler,
)

var serviceSet = wire.NewSet(
	imp.NewChatService,
)

var repositorySet = wire.NewSet(
	chat.NewChatCommandRepository,
	chat.NewChatQueryRepository,
)

func InitializeServer(
	cfg *config.Config,
	db *database.MongoDbClient,
	rdb *redis.RedisClient,
	userClient grpcUser.UserServiceClient,
	commentWS *ws.Socket,
) *api.Server {
	wire.Build(serverSet, itemServerSet, httpHandlerSet, serviceSet, repositorySet)
	return &api.Server{}
}
