package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.\

import (
	usecase "github.com/wrferreira1003/Desafio-Clean-Architecture/internal/usecase"
)

// Resolver é a implementação do resolver root
type Resolver struct {
	OrderUseCase usecase.OrderUseCaseInterface
}
