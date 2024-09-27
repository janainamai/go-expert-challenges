package resources

import (
	"github.com/janainamai/go-expert-challenges/cmd/configs"
	"github.com/janainamai/go-expert-challenges/internal/entity/usecase"
	"github.com/janainamai/go-expert-challenges/internal/entity/usecase/gateway"
	"github.com/janainamai/go-expert-challenges/internal/infra/mysql"
	setupMySQL "github.com/janainamai/go-expert-challenges/internal/infra/mysql/setup"
	"github.com/janainamai/go-expert-challenges/internal/infra/rest"
)

type (
	Resources struct {
		SaveOrderRestHandler  rest.SaveOrderRestHandler
		ListOrdersRestHandler rest.ListOrdersRestHandler
	}
)

func LoadResources(cfg *configs.Config) *Resources {

	// infra mysql
	mysqlConnector := setupMySQL.NewMySQLConnector(cfg)
	saveOrderMySQL := mysql.NewSaveOrderMySQL(mysqlConnector)
	listOrdersMySQL := mysql.NewListOrdersMySQL(mysqlConnector)

	// gateway
	saveOrderGateway := gateway.NewSaveOrderGateway(saveOrderMySQL)
	listOrdersGateway := gateway.NewListOrdersGateway(listOrdersMySQL)

	// usecase
	saveOrderUseCase := usecase.NewSaveOrderUseCase(saveOrderGateway)
	listOrdersUseCase := usecase.NewListOrdersUseCase(listOrdersGateway)

	// infra rest
	saveOrderRest := rest.NewSaveOrderRestHandler(saveOrderUseCase)
	listOrdersRest := rest.NewListOrdersRestHandler(listOrdersUseCase)

	return &Resources{
		SaveOrderRestHandler:  *saveOrderRest,
		ListOrdersRestHandler: *listOrdersRest,
	}
}
