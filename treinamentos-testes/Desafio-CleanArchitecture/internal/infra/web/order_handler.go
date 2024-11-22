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
