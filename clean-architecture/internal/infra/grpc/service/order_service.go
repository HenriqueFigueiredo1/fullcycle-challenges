package service

import (
	"context"

	"github.com/HenriqueFigueiredo1/clean-architecture/internal/dto"
	"github.com/HenriqueFigueiredo1/clean-architecture/internal/infra/grpc/pb"
	"github.com/HenriqueFigueiredo1/clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	listOrderUseCase usecase.ListOrderUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(_ context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	orderDto := dto.OrderInput{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(orderDto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrder(_ context.Context, _ *pb.ListOrderRequest) (*pb.ListOrderResponse, error) {
	outputOrders, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	orderList := &pb.ListOrderResponse{
		Orders: make([]*pb.Order, 0),
	}

	for _, order := range outputOrders {
		orderList.Orders = append(orderList.Orders, &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}

	return orderList, nil
}
