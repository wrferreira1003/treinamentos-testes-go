package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/wrferreira1003/api-server-go/internal/dto"
	"github.com/wrferreira1003/api-server-go/internal/entity"
	"github.com/wrferreira1003/api-server-go/internal/infra/database"
	"github.com/wrferreira1003/api-server-go/pkg/utils"
)

// Criar os Handler, pois eles são responsáveis por receber as requisições e retornar as respostas
type ProductHandler struct {
	ProductDB database.ProductDB
}

// Esse método é responsável por criar o handler, handler é o responsável por receber a requisição e retornar a resposta
func NewProductHandler(db database.ProductDB) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create Product godoc
// @Summary Create Product
// @Description Create Product
// @Tags Products
// @Accept json
// @Produce json
// @Param request body dto.CreateProductInput true "use request"
// @Success 201
// @Failure 500 {object} Error
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	//Passando pelo DTO para evitar acesso direto na entidade, recebendo os dados conforme o dto, caso esteja errado, retorna erros
	var product dto.CreateProductInput // Criando uma variável do tipo CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Esse acesso direto na entidade aqui no handler é errado,
	//pois o handler é responsável por receber a requisição e retornar a resposta, e não acessar a entidade diretamente
	//Normalmente fazemos isso pelo caso de uso.
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Salvando o produto no banco de dados
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Get Product godoc
// @Summary Get Product
// @Description Get Product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" Format(uuid)
// @Success 200 {object} entity.Product
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// Verificando se o id está vazio
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Buscando o produto no banco de dados
	p, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Retornando o produto no formato JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

// Update Product godoc
// @Summary Update Product
// @Description Update Product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" Format(uuid)
// @Param request body dto.CreateProductInput true "product update request"
// @Success 200
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Buscando o produto no banco de dados
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Validando o ID
	product.ID, err = utils.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verificando se o produto existe
	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Atualizando o produto no banco de dados
	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete Product godoc
// @Summary Delete Product
// @Description Delete Product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" Format(uuid)
// @Success 200
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Buscar o produto pelo ID
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Deletar o produto passando o ponteiro
	err = h.ProductDB.Delete(product) // Passando o ponteiro do produto encontrado
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Get Product godoc
// @Summary Get Product
// @Description Get Product
// @Tags Products
// @Accept json
// @Produce json
// @Param page query string false "Page number"
// @Param limit query string false "Number of products per page"
// @Param sort query string false "Sort by"
// @Success 200 {array} entity.Product
// @Failure 500 {object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	// convertendo para int
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
