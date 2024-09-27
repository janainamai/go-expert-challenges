package service

import (
	"context"
	"fmt"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/domain"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/usecase"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/grpc/pb"
)

type orderService struct {
	pb.UnimplementedOrderServiceServer // implementation for not implemented funcs
	createOrderUsecase                 usecase.SaveOrderUseCaseInterface
	listOrdersUsecase                  usecase.ListOrdersUseCaseInterface
}

func NewOrdeGrpcService(createOrderUsecase usecase.SaveOrderUseCaseInterface, listOrdersUsecase usecase.ListOrdersUseCaseInterface) pb.OrderServiceServer {
	return &orderService{
		createOrderUsecase: createOrderUsecase,
		listOrdersUsecase:  listOrdersUsecase,
	}
}

func (c *orderService) CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*pb.Order, error) {
	domain := &domain.Order{
		Price: float64(request.Price),
		Tax:   float64(request.Tax),
	}

	order, err := c.createOrderUsecase.Save(ctx, domain)
	if err != nil {
		return nil, fmt.Errorf("error: %s, detail: %s", err.Title, err.Detail)
	}

	return &pb.Order{
		Id:         order.ID,
		Price:      float32(order.Price),
		Tax:        float32(order.Tax),
		FinalPrice: float32(order.FinalPrice),
	}, nil
}

func (c *orderService) ListOrders(ctx context.Context, request *pb.Blank) (*pb.OrdersList, error) {
	ordersDomain, err := c.listOrdersUsecase.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error: %s, detail: %s", err.Title, err.Detail)
	}

	var orders []*pb.Order
	for _, domain := range ordersDomain {
		order := &pb.Order{
			Id:         domain.ID,
			Price:      float32(domain.Price),
			Tax:        float32(domain.Tax),
			FinalPrice: float32(domain.FinalPrice),
		}

		orders = append(orders, order)
	}

	return &pb.OrdersList{Orders: orders}, nil
}
