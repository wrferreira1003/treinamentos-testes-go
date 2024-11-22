package usecase

import (
	dto "github.com/wrferreira1003/Desafio-Clean-Architecture/internal/Dto"
	entities "github.com/wrferreira1003/Desafio-Clean-Architecture/internal/domain/entities"
	repository "github.com/wrferreira1003/Desafio-Clean-Architecture/internal/domain/repository"
	events "github.com/wrferreira1003/Desafio-Clean-Architecture/pkg/events"
)

type CreateOrderUseCase struct {
	OrderRepository repository.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateOrderUseCase(
	orderRepository repository.OrderRepositoryInterface,
	orderCreated events.EventInterface,
	eventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: orderRepository,
		OrderCreated:    orderCreated,
		EventDispatcher: eventDispatcher,
	}
}

func (c *CreateOrderUseCase) Execute(input dto.OrderInputDTO) (dto.OrderOutputDTO, error) {
	order, err := entities.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return dto.OrderOutputDTO{}, err
	}

	// Calcula o preço final do pedido e salva no banco de dados caso não ocorra erro
	order.CalculateFinalPrice()
	if err := c.OrderRepository.Save(order); err != nil {
		return dto.OrderOutputDTO{}, err
	}

	output := dto.OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	c.OrderCreated.GetPayload()                                      // Define os dados que estao sendo gerados no evento
	c.EventDispatcher.Dispatch("clean_architecture", c.OrderCreated) // Dispara o evento

	return output, nil
}
