PROTO_DIR = internal/infra/grpc/proto

.PHONY: protoc
protoc-gen:
	protoc --go_out=. --go-grpc_out=. $(PROTO_DIR)/entities.proto