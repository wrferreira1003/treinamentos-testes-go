package product

type ProductUsecase struct {
	productRepository ProductInterface
}

func NewProductUsecase(productRepository ProductInterface) *ProductUsecase {
	return &ProductUsecase{productRepository: productRepository}
}

// Criando um metodo para buscar um produto pelo ID, Lembrando que o usecase nao deve retornar a entidade, mas um DTO
func (u *ProductUsecase) GetProduct(id int) (Product, error) {
	return u.productRepository.GetProduct(id)
}
