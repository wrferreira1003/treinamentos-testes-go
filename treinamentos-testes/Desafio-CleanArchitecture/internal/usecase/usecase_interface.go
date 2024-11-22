package usecase

import dto "github.com/wrferreira1003/Desafio-Clean-Architecture/internal/Dto"

type OrderUseCaseInterface interface {
	Execute(input dto.OrderInputDTO) (dto.OrderOutputDTO, error)
}
