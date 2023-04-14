package repository

import (
	"context"

	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/domain"
	interfaces "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/repository/interface"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

// GetProduct implements interfaces.ProductRepository
func (p *productDatabase) GetProduct(ctx context.Context, id int64) (domain.Product, error) {
	var product domain.Product
	err := p.DB.Where("id =", id).First(&product).Error
	return product, err
}

// CreateProduct implements interfaces.ProductRepository
func (p *productDatabase) CreateProduct(ctx context.Context, product domain.Product) (int64, error) {
	err := p.DB.Create(&product).Error
	return product.Id, err
}

func NewProductDatabase(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{
		DB,
	}
}
