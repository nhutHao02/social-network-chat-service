run:
	go run main.go --config=./config/local/config.yaml
wire:
	wire ./internal
proto-gen:
	protoc -I=pkg/grpc/proto --go_out=pkg/grpc --go_opt=paths=source_relative \
    --go-grpc_out=pkg/grpc --go-grpc_opt=paths=source_relative \
    pkg/grpc/proto/*.proto