package entity

import (
	"errors"
	"time"

	"github.com/wrferreira1003/api-server-go/pkg/utils"
)

// Criar as variaveis de erro
var (
	errIDIsRequired = errors.New("id is required")
	errIDIsInvalid  = errors.New("id is invalid")

	ErrNameIsRequired = errors.New("name is required")

	ErrPriceIsRequired = errors.New("price is required")
	ErrPriceIsInvalid  = errors.New("price is invalid")
)

// Product represents a product entity
type Product struct {
	ID        utils.ID  `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	// Cria o produto
	product := &Product{
		ID:        utils.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	// Validação do produto
	if err := product.Validate(); err != nil {
		return nil, err
	}
	// Retorna o produto
	return product, nil
}

// Validação do produto
func (p *Product) Validate() error {
	// Validação do ID
	if p.ID.String() == "" {
		return errIDIsRequired
	}
	// verifico se o ID é válido
	if _, err := utils.ParseID(p.ID.String()); err != nil {
		return errIDIsInvalid
	}
	// Validação do Name
	if p.Name == "" {
		return ErrNameIsRequired
	}
	// Validação do Price
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	// Verifica se o preço é válido
	if p.Price < 0 {
		return ErrPriceIsInvalid
	}
	return nil
}
