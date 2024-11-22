//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/domain/repository"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/infra/database"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/usecase"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/pkg/events"
)

// set de provedores seria tudo que vamos utilizar no projeto
var OrderSet = wire.NewSet(
	// Registro de provedores
	database.NewOrderRepository,
	// Registro de binds para interfaces e suas implementações concretas
	wire.Bind(new(repository.OrderRepositoryInterface), new(*database.OrderRepository)),

	events.NewEventDispatcher,
	// Registro de binds para interfaces e suas implementações concretas
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),

	events.NewOrderCreatedEvent,

	// Registro de usecases
	usecase.NewCreateOrderUseCase,
)

// Inicializador gerado pelo Google Wire para o OrderUseCase
func InitializeOrderUseCase(db *sql.DB) (*usecase.CreateOrderUseCase, error) {
	wire.Build(OrderSet)
	return &usecase.CreateOrderUseCase{}, nil
}
