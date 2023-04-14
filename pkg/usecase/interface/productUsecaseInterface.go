package interfaces

import (
	"context"

	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/domain"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, product domain.Product) (int64, error)
	GetProduct(ctx context.Context, id int64) (domain.Product, error)
	ListProducts(ctx context.Context) ([]domain.Product, error)
	UpdateProduct(ctx context.Context, product domain.Product) (int64, error)
	DeleteProduct(ctx context.Context, id int64) error
}
