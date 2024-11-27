package startup

import (
	"context"
	"fmt"
	"log"

	"github.com/nhutHao02/social-network-chat-service/config"
	"github.com/nhutHao02/social-network-chat-service/database"
	"github.com/nhutHao02/social-network-chat-service/internal"
	"github.com/nhutHao02/social-network-chat-service/internal/api"
	"github.com/nhutHao02/social-network-chat-service/pkg/redis"
	"github.com/nhutHao02/social-network-chat-service/pkg/websocket"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	pb "github.com/nhutHao02/social-network-user-service/pkg/grpc"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Start() {
	// init logger
	initLogger()

	// load congig
	cfg := config.LoadConfig()

	// database setup
	db := database.ConnectToMongo(cfg.Database)

	// init redis
	rdb := redis.InitRedis(cfg.Redis)
	// Test connection
	_, err := rdb.Rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("failed to init redis------------", zap.Error(err))
		panic(fmt.Sprintf("Could not connect to Redis: %v", err))
	}

	// connect to grpc server
	userClient := openClientConnection(cfg.Client)

	// init Socket
	ws := websocket.NewSocket()

	// // init Server
	server := internal.InitializeServer(cfg, db, rdb, userClient, ws)

	// run server
	runServer(server)

}

func runServer(server *api.Server) {
	var g errgroup.Group

	g.Go(func() error {
		return server.HTTPServer.RunHTTPServer()
	})

	// g.Go(func() error {
	// 	return server.GRPCServer.RunGRPCServer()
	// })

	if err := g.Wait(); err != nil {
		logger.Fatal("Error when start server", zap.Error(err))
	}
}

func initLogger() {
	err := logger.InitLogger()
	if err != nil {
		log.Fatalf("Could not initialize logger: %v", err)
	}
	defer logger.Sync()
}

func openClientConnection(cfg *config.ClientConfig) pb.UserServiceClient {
	conn, err := grpc.NewClient(cfg.UserService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("Failed to connect to gRPC server: ", zap.Error(err))
	}
	client := pb.NewUserServiceClient(conn)
	logger.Info("Connect to user gRPC server port: " + cfg.UserService)
	return client
}
