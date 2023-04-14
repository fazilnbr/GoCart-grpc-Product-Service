package services

import (
	usecase "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/usecase/interface"
)

type ProductService struct {
	userUseCase usecase.ProductUseCase
}

func NewProductService(usecase usecase.ProductUseCase) *ProductService {
	return &ProductService{
		userUseCase: usecase,
	}
}
