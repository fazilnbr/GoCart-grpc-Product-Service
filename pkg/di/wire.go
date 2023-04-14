//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/api"
	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/api/handler"
	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/config"
	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/db"
	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/repository"
	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
