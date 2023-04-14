package usecase

import (
	"context"

	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/domain"
	repository "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/repository/interface"
	interfaces "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/usecase/interface"
)

type productUseCase struct {
	userRepo repository.ProductRepository
}

// GetProduct implements interfaces.ProductUseCase
func (p *productUseCase) GetProduct(ctx context.Context, id int64) (domain.Product, error) {
	product,err:=p.userRepo.GetProduct(ctx,id)
	return product,err
}

// CreateProduct implements interfaces.ProductUseCase
func (p *productUseCase) CreateProduct(ctx context.Context, product domain.Product) (int64, error) {
	id, err := p.userRepo.CreateProduct(ctx, product)
	return id, err
}

func NewproductUseCase(repo repository.ProductRepository) interfaces.ProductUseCase {
	return &productUseCase{
		userRepo: repo,
	}
}
