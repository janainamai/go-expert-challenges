package graph

import "github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	SaveOrderUseCase  usecase.SaveOrderUseCaseInterface
	ListOrdersUseCase usecase.ListOrdersUseCaseInterface
}
