package database

import (
	"github.com/wrferreira1003/api-server-go/internal/entity"
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

// Criar uma instância de ProductDB
func NewProduct(db *gorm.DB) *ProductDB {
	return &ProductDB{DB: db}
}

// Cria um novo produto
func (p *ProductDB) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

// Busca um produto por ID
func (p *ProductDB) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Atualiza um produto
func (p *ProductDB) Update(product *entity.Product) error {
	// Verifica se o produto existe
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	// Atualiza o produto
	return p.DB.Save(product).Error
}

// Deleta um produto
func (p *ProductDB) Delete(product *entity.Product) error {
	// Verifica se o produto existe
	productFound, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	// Deleta o produto
	return p.DB.Delete(productFound).Error
}

// Lista todos os produtos com paginação
func (p *ProductDB) FindAll(page int, limit int, sort string) ([]entity.Product, error) {
	products := []entity.Product{}
	var err error
	//valida se a paginação é valida
	if sort != "asc" && sort != "desc" && sort != "" {
		sort = "asc"
	}
	if page != 0 || limit != 0 {
		// Verifica se a paginação é valida
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}
	return products, err
}
