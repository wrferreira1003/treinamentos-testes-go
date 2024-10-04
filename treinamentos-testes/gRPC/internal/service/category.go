package service

import (
	"context"

	"github.com/rcfacil/gRPC-go/internal/database"
	"github.com/rcfacil/gRPC-go/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCourseCategoryServer
	categoryDB *database.Category
}

// Metodo construtor
func NewCategoryService(categoryDB *database.Category) *CategoryService {
	return &CategoryService{categoryDB: categoryDB}
}

// Cria uma nova categoria
func (c *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	// Cria a categoria no banco de dados
	category, err := c.categoryDB.Create(req.Name, req.Description)
	if err != nil {
		return nil, err
	}
	// Converte a categoria para o formato do protobuf
	categoryResponse := &pb.CategoryResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}

	return categoryResponse, nil
}
