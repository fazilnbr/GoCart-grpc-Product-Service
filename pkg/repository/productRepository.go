package repository

import (
	interfaces "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/repository/interface"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

// // CreateProduct implements interfaces.UserRepository
// func (*userDatabase) CreateProduct(ctx context.Context, user domain.Product) (int, error) {

// }

func NewProductDatabase(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB}
}
