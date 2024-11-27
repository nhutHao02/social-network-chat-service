package chat

type ChatQueryRepository interface {
	ChatQuery()
}

type ChatCommandRepository interface {
	ChatCommand()
}
