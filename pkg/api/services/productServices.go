package services

import (
	"context"
	"net/http"

	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/domain"
	"github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/pb"
	usecase "github.com/fazilnbr/GoCart-grpc-Product-Service/pkg/usecase/interface"
)

type ProductService struct {
	userUseCase usecase.ProductUseCase
}

func (p *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product:=domain.Product{
		Id: req.Id,
		Name: req.Name,
		Description: req.Description,
		Price: float64(req.Price),
		Stock: req.Stock,
	}
	id,err:=p.userUseCase.CreateProduct(ctx,product)

	if err!=nil{
		return &pb.CreateProductResponse{
			Status: http.StatusUnprocessableEntity,
			Error: err.Error(),
		},err
	}

	return &pb.CreateProductResponse{
		Status: http.StatusOK,
		Id: id,
	},nil


}
func NewProductService(usecase usecase.ProductUseCase) *ProductService {
	return &ProductService{
		userUseCase: usecase,
	}
}
