package services

import (
	"context"

	dto "github.com/wrferreira1003/Desafio-Clean-Architecture/internal/Dto"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/infra/grpc/pb"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.OrderUseCaseInterface
}

func NewOrderService(createOrderUseCase usecase.OrderUseCaseInterface) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := dto.OrderInputDTO{
		ID:    req.Id,
		Price: float64(req.Price),
		Tax:   float64(req.Tax),
	}

	// Executa o caso de uso para criar o pedido
	output, err := o.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	// Retorna a resposta ao cliente
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}
