package handler

import (
	usecase "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}
