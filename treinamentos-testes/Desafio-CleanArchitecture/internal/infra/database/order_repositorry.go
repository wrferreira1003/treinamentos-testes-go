package database

import (
	"database/sql"

	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/domain/entities"
)

type OrderRepository struct {
	Database *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Database: db,
	}
}

func (r *OrderRepository) Save(order *entities.Order) error {
	// Prepara e bom quando precisamos executar a mesma query varias vezes
	stmt, err := r.Database.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	// Executa a query
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	defer stmt.Close()
	return err
}
