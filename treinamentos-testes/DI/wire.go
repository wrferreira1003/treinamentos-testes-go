//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/wrferreira1003/DI/product"
)

var setRepository = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductInterface), new(*product.ProductRepository)),
)

// Criando uma função para criar uma instancia do usecase
func NewUsecase(db *sql.DB) *product.ProductUsecase {
	wire.Build(
		product.NewProductUsecase,
		setRepository,
	)
	return &product.ProductUsecase{}
}
