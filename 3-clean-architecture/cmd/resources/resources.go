package resources

import (
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/cmd/configs"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/usecase"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/usecase/gateway"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/grpc/pb"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/grpc/service"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/mysql"
	setupMySQL "github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/mysql/setup"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/rest"
)

type (
	Resources struct {
		CreateOrderRestHandler rest.CreateOrderRestHandler
		ListOrdersRestHandler  rest.ListOrdersRestHandler

		OrderGrpcService pb.OrderServiceServer

		CreateOrderUseCase usecase.CreateOrderUseCaseInterface
		ListOrdersUseCase  usecase.ListOrdersUseCaseInterface
	}
)

func LoadResources(cfg *configs.Config) *Resources {

	// infra out: mysql
	mysqlConnector := setupMySQL.NewMySQLConnector(cfg)
	saveOrderMySQL := mysql.NewSaveOrderMySQL(mysqlConnector)
	listOrdersMySQL := mysql.NewListOrdersMySQL(mysqlConnector)

	// gateway usecase
	createOrderGateway := gateway.NewCreateOrderGateway(saveOrderMySQL)
	listOrdersGateway := gateway.NewListOrdersGateway(listOrdersMySQL)

	// usecase
	createOrderUseCase := usecase.NewCreateOrderUseCase(createOrderGateway)
	listOrdersUseCase := usecase.NewListOrdersUseCase(listOrdersGateway)

	// infra in: rest
	createOrderRest := rest.NewCreateOrderRestHandler(createOrderUseCase)
	listOrdersRest := rest.NewListOrdersRestHandler(listOrdersUseCase)

	// infra in: grpc
	orderGrpc := service.NewOrdeGrpcService(createOrderUseCase, listOrdersUseCase)

	return &Resources{
		CreateOrderRestHandler: *createOrderRest,
		ListOrdersRestHandler:  *listOrdersRest,

		OrderGrpcService: orderGrpc,

		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}
