package repository

import (
	"context"
	"errors"

	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/domain"
	interfaces "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/repository/interface"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

// DeleteProduct implements interfaces.ProductRepository
func (p *productDatabase) DeleteProduct(ctx context.Context, id int64) error {
	var product domain.Product
	return p.DB.Where("id = ?", id).Delete(&product).Error
}

// UpdateProduct implements interfaces.ProductRepository
func (p *productDatabase) UpdateProduct(ctx context.Context, product domain.Product) (int64, error) {
	result := p.DB.Save(&product)
	return product.Id, result.Error
}

// ListProducts implements interfaces.ProductRepository
func (p *productDatabase) ListProducts(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	result := p.DB.Find(&products)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("there is no product to desplay")
	}
	return products, result.Error

}

// GetProduct implements interfaces.ProductRepository
func (p *productDatabase) GetProduct(ctx context.Context, id int64) (domain.Product, error) {
	var product domain.Product
	err := p.DB.Where("id = ?", id).First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Product{}, errors.New("there is no product to desplay on this id")
	}
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
