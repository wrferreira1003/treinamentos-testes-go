package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wrferreira1003/api-server-go/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// criando produto para teste
func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	// Migrate the schema
	db.AutoMigrate(&entity.Product{}) // criando tabela de produtos

	product, err := entity.NewProduct("Product 1", 1000)
	assert.NoError(t, err)    // verifica se o erro é nulo
	assert.NotNil(t, product) // verifica se o produto não é nulo

	productDb := NewProduct(db)
	err = productDb.Create(product) // criando produto no banco de dados
	assert.NoError(t, err)          // verifica se o erro é nulo
	assert.NotNil(t, product)       // verifica se o produto não é nulo
	assert.NotEmpty(t, product.ID)  // verifica se o ID do produto não é vazio

}

// Testando a busca com paginação
func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{}) // criando tabela de produtos

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100) // criando produto
		assert.NoError(t, err)                                                              // verifica se o erro é nulo
		db.Create(product)                                                                  // criando produto no banco de dados                                                         // verifica se o erro é nulo
	}

	productDB := NewProduct(db)
	product, err := productDB.FindAll(1, 10, "asc")                                                 // buscando todos os produtos
	assert.NoError(t, err, "Failed to find products")                                               // verifica se o erro é nulo
	assert.Len(t, product, 10, "Expected 10 products on page 1")                                    // verifica se o tamanho do produto é 10
	assert.Equal(t, "Product 1", product[0].Name, "First product on page 1 should be 'Product 1'")  // verifica se o nome do produto é igual a "Product 1"
	assert.Equal(t, "Product 10", product[9].Name, "Last product on page 1 should be 'Product 10'") // verifica se o nome do produto é igual a "Product 10"

	product, err = productDB.FindAll(2, 10, "asc")                                                   // buscando todos os produtos
	assert.NoError(t, err, "Failed to find products on page 2")                                      // verifica se o erro é nulo
	assert.Len(t, product, 10, "Expected 10 products on page 2")                                     // verifica se o tamanho do produto é 10
	assert.Equal(t, "Product 11", product[0].Name, "First product on page 2 should be 'Product 11'") // verifica se o nome do produto é igual a "Product 11"
	assert.Equal(t, "Product 20", product[9].Name, "Last product on page 2 should be 'Product 20'")  // verifica se o nome do produto é igual a "Product 20"

	product, err = productDB.FindAll(3, 10, "asc")                                                   // buscando todos os produtos
	assert.NoError(t, err, "Failed to find products on page 3")                                      // verifica se o erro é nulo
	assert.Len(t, product, 3, "Expected 3 products on page 3")                                       // verifica se o tamanho do produto é 3
	assert.Equal(t, "Product 21", product[0].Name, "First product on page 3 should be 'Product 21'") // verifica se o nome do produto é igual a "Product 21"
	assert.Equal(t, "Product 23", product[2].Name, "Last product on page 3 should be 'Product 23'")  // verifica se o nome do produto é igual a "Product 23"

}

// Find by ID
func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{}) // creating product table

	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)    // check if error is nil
	assert.NotNil(t, product) // check if product is not nil

	productDB := NewProduct(db) // creating product instance

	// Save the product to the database
	err = db.Create(&product).Error
	assert.NoError(t, err) // check if error is nil

	// Retrieve the product by ID
	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)                  // check if error is nil
	assert.NotNil(t, product)               // check if product is not nil
	assert.Equal(t, product.ID, product.ID) // check if the product ID matches
}

// update product
func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{}) // creating product table

	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)    // check if error is nil
	assert.NotNil(t, product) // check if product is not nil

	// Save the product to the database
	err = db.Create(&product).Error
	assert.NoError(t, err) // check if error is nil

	productDB := NewProduct(db) // creating product instance

	// Update the product
	product.Name = "Product 1 Updated"
	err = productDB.Update(product)                        // updating product
	assert.NoError(t, err)                                 // check if error is nil
	product, err = productDB.FindByID(product.ID.String()) // finding product by ID
	assert.NoError(t, err)                                 // check if error is nil
	assert.Equal(t, product.Name, "Product 1 Updated")     // check if the product name is updated
}

// Delete product
func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{}) // creating product table

	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)    // check if error is nil
	assert.NotNil(t, product) // check if product is not nil

	db.Create(&product)

	productDB := NewProduct(db) // creating product instance

	err = productDB.Delete(product) // deleting product
	assert.NoError(t, err)          // check if error is nil

	product, err = productDB.FindByID(product.ID.String()) // finding product by ID
	assert.Error(t, err)                                   // check if error is not nil
	assert.Nil(t, product)                                 // check if product is nil

}
