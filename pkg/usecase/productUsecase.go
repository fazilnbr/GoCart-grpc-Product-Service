package usecase

import (
	repository "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/repository/interface"
	interfaces "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/usecase/interface"
)

type productUseCase struct {
	userRepo repository.ProductRepository
}

func NewproductUseCase(repo repository.ProductRepository) interfaces.ProductUseCase {
	return &productUseCase{
		userRepo: repo,
	}
}
