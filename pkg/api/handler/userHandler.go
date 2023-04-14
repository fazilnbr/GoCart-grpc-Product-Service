package handler

import (
	usecase "github.com/fazilnbr/banking-grpc-account-service/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}
