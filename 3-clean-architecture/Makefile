PROTO_DIR = internal/infra/grpc/proto

.PHONY: protoc

grpc-gen:
	protoc --go_out=. --go-grpc_out=. $(PROTO_DIR)/entities.proto

graph-gen:
	go run github.com/99designs/gqlgen generate