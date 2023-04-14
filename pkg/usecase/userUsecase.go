package usecase

import (
	repository "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/repository/interface"
	interfaces "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) interfaces.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}
