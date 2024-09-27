package setup

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/cmd/configs"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/cmd/resources"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/graph"
)

func InitServer(cfg *configs.Config, resources *resources.Resources) {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		SaveOrderUseCase:  resources.SaveOrderUseCase,
		ListOrdersUseCase: resources.ListOrdersUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Printf("GraphQL Server - Listening on port: %s\n", cfg.GraphServer.Port)
	if err := http.ListenAndServe(":"+cfg.GraphServer.Port, nil); err != nil {
		panic(fmt.Sprintf("error initing graphql server: %s", err.Error()))
	}
}
