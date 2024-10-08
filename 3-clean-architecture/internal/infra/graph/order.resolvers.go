package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"
	"fmt"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/domain"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/graph/model"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*model.Order, error) {
	domain := &domain.Order{
		Price: input.Price,
		Tax:   input.Tax,
	}

	order, err := r.CreateOrderUseCase.Create(ctx, domain)
	if err != nil {
		return nil, fmt.Errorf("error: %s, detail: %s", err.Title, err.Detail)
	}

	return &model.Order{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	ordersDomain, err := r.ListOrdersUseCase.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error: %s, detail: %s", err.Title, err.Detail)
	}

	var orders []*model.Order
	for _, domain := range ordersDomain {
		order := &model.Order{
			ID:         domain.ID,
			Price:      domain.Price,
			Tax:        domain.Tax,
			FinalPrice: domain.FinalPrice,
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
