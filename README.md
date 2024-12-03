# social-network-chat-service
## Project Summary
This is project about social network that allows users to share content, images, and emotions, and have real-time communication capabilities, while ensuring high performance, security, and scalability using the microservices architecture.

#### Technologies:
- Back-end:
  - Language: Go.
  - Frameworks/Platforms: Gin-Gonic, gRPC, Swagger, JWT, Google-Wire, SQLX, Redis, Zap, WebSocket.
  - Database: MariaDB, MongoDB.
- Front-end:
  - Language: JavaScript.
  - Frameworks/Platforms: React, Tailwind CSS, FireBase.

## The project includes repositories
- [common-service](https://github.com/nhutHao02/social-network-common-service)
- [user-service](https://github.com/nhutHao02/social-network-user-service)
- [tweet-service](https://github.com/nhutHao02/social-network-tweet-service)
- [chat-service](https://github.com/nhutHao02/social-network-chat-service)
- [notification-service](https://github.com/nhutHao02/social-network-notification-service)
- [Front-end-service (in progress)](https://github.com/nhutHao02/)

## This service
This is the service that provides the APIs related to the Chat and handle the real-time messages.

## Project structure
```
.
├── config
│   ├── config.go
│   └── local
│       └── config.yaml
├── database
│   └── database.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── grpc
│   │   ├── http
│   │   │   ├── http_server.go
│   │   │   └── v1
│   │   │       ├── chat_handler.go
│   │   │       └── route.go
│   │   └── server.go
│   ├── application
│   │   ├── chat_service.go
│   │   └── imp
│   │       └── chat_service_imp.go
│   ├── domain
│   │   ├── entity
│   │   │   ├── chat.go
│   │   │   └── collection_name.go
│   │   ├── interface
│   │   │   └── chat
│   │   │       └── chat_repository.go
│   │   └── model
│   │       ├── chat.go
│   │       ├── user.go
│   │       └── websocket.go
│   ├── infrastructure
│   │   └── chat
│   │       ├── command_repository.go
│   │       └── query_repository.go
│   ├── wire_gen.go
│   └── wire.go
├── main.go
├── Makefile
├── pkg
│   ├── common
│   │   └── response.go
│   ├── constants
│   │   └── constants.go
│   ├── redis
│   │   └── redis.go
│   └── websocket
│       └── websocket.go
├── README.md
└── startup
    └── startup.go
```