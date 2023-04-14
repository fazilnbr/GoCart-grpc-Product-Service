package interfaces

import (
	"context"

	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/domain"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product domain.Product) (int64, error)
}
