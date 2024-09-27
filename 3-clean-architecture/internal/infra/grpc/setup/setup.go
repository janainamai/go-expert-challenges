package setup

import (
	"fmt"
	"net"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/cmd/configs"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/cmd/resources"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InitService(cfg *configs.Config, resources *resources.Resources) {

	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, resources.OrderGrpcService)
	reflection.Register(server) // to work with Evans

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcService.Port))
	if err != nil {
		panic(err)
	}

	fmt.Printf("gRPC Server - Listening on port: %s\n", cfg.GrpcService.Port)
	if err := server.Serve(lis); err != nil {
		panic(fmt.Sprintf("error initing grpc server: %s", err.Error()))
	}

}
