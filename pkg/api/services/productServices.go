package services

import (
	usecase "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/usecase/interface"
)

type ProductService struct {
	userUseCase usecase.UserUseCase
}

func NewProductService(usecase usecase.UserUseCase) *ProductService {
	return &ProductService{
		userUseCase: usecase,
	}
}
