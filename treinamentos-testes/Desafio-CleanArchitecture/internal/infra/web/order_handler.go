package web

import (
	"encoding/json"
	"net/http"

	dto "github.com/wrferreira1003/Desafio-Clean-Architecture/internal/Dto"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/usecase"
)

type WebOrderHandler struct {
	CreateOrderUseCase usecase.OrderUseCaseInterface
}

func NewWebOrderHandler(
	createOrderUseCase usecase.OrderUseCaseInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dto.OrderInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	/* Se alguem tiver corrigindo esse codigo, pode comentar no retorno do trabalho, porque o Wesley
	Iniciou o usecase dentro do handler, onde fere os principios de responsabilidade unica e de inversao de controle
	na minha humilde opiniao. Criei uma interface para o usecase para facilitar, no meu entendimento, dessa forma facilita
	a criacao de testes unitarios, além de poder ser reutilizado em outras partes do sistema.
	*/

	output, err := h.CreateOrderUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
