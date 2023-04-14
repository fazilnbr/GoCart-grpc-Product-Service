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

// CreateProduct implements interfaces.ProductRepository
func (c *productDatabase) CreateProduct(ctx context.Context, product domain.Product) (int64, error) {
	err := c.DB.Create(&product).Error
	return product.Id, err
}

func NewProductDatabase(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{
		DB,
	}
}
