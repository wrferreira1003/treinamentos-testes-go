package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Teste para verificar se o produto é criado corretamente
func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 10.0) // Cria o produto
	assert.Nil(t, err)                            // Verifica se não houve erro
	assert.NotNil(t, product)                     // Verifica se o produto não é nulo
	assert.NotEmpty(t, product.ID)                // Verifica se o ID do produto não é vazio
	assert.Equal(t, "Product 1", product.Name)    // Verifica se o nome do produto é igual ao esperado
	assert.Equal(t, 10.0, product.Price)          // Verifica se o preço do produto é igual ao esperado
}

// Teste para verificar se o nome é obrigatório
func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 10.0)    // Cria o produto
	assert.Nil(t, product)                  // Verifica se o produto é nulo
	assert.Equal(t, ErrNameIsRequired, err) // Verifica se o erro é igual ao esperado
}

// Teste para verificar se o preço é obrigatório
func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0) // Cria o produto
	assert.Nil(t, product)                     // Verifica se o produto é nulo
	assert.Equal(t, ErrPriceIsRequired, err)   // Verifica se o erro é igual ao esperado
}

// Teste para verificar se o preço é inválido
func TestProductWhenInvalidPrice(t *testing.T) {
	product, err := NewProduct("Product 1", -10.0) // Cria o produto
	assert.Nil(t, product)                         // Verifica se o produto é nulo
	assert.Equal(t, ErrPriceIsInvalid, err)        // Verifica se o erro é igual ao esperado
}

// Teste para verificar se o produto é válido
func TestProductValidate(t *testing.T) {
	product, err := NewProduct("Product 1", 10.0) // Cria o produto
	assert.Nil(t, err)                            // Verifica se não houve erro
	assert.NotNil(t, product)                     // Verifica se o produto não é nulo
	assert.NotEmpty(t, product.ID)                // Verifica se o ID do produto não é vazio
	assert.Nil(t, product.Validate())             // Verifica se o produto é válido
}
