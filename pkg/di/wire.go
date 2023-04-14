//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/api"
	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/api/services"
	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/config"
	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/db"
	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/repository"
	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewProductDatabase,
		usecase.NewproductUseCase,
		services.NewProductService,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
