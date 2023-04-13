//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/fazilnbr/banking-grpc-account-service/pkg/api"
	"github.com/fazilnbr/banking-grpc-account-service/pkg/api/handler"
	"github.com/fazilnbr/banking-grpc-account-service/pkg/config"
	"github.com/fazilnbr/banking-grpc-account-service/pkg/db"
	"github.com/fazilnbr/banking-grpc-account-service/pkg/repository"
	"github.com/fazilnbr/banking-grpc-account-service/pkg/usecase"
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
