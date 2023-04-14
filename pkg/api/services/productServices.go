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

func (p *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	err := p.userUseCase.DeleteProduct(ctx, req.Id)
	if err != nil {
		return &pb.DeleteProductResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

	return &pb.DeleteProductResponse{
		Status: http.StatusOK,
		Id:     req.Id,
	}, nil
}

func (p *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	product := domain.Product{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Price:       float32(req.Price),
		Stock:       req.Stock,
	}
	id, err := p.userUseCase.UpdateProduct(ctx, product)

	if err != nil {
		return &pb.UpdateProductResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

	return &pb.UpdateProductResponse{
		Status: http.StatusOK,
		Id:     id,
	}, nil

}

func (p *ProductService) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductResponse, error) {

	products, err := p.userUseCase.ListProducts(ctx)

	if err != nil {
		return &pb.ListProductResponse{
			Status:   http.StatusUnprocessableEntity,
			Error:    err.Error(),
			Products: nil,
		}, err
	}

	var pbProducts []*pb.Product
	for _, p := range products {
		pbProducts = append(pbProducts, &pb.Product{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
		})
	}
	return &pb.ListProductResponse{
		Status:   http.StatusOK,
		Products: pbProducts,
	}, err

}

func (p *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, err := p.userUseCase.GetProduct(ctx, req.Id)
	if err != nil {
		return &pb.GetProductResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

	return &pb.GetProductResponse{
		Status:      http.StatusOK,
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       float32(product.Price),
		Stock:       product.Stock,
	}, nil

}

func (p *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := domain.Product{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Price:       float32(req.Price),
		Stock:       req.Stock,
	}
	id, err := p.userUseCase.CreateProduct(ctx, product)

	if err != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

	return &pb.CreateProductResponse{
		Status: http.StatusOK,
		Id:     id,
	}, nil

}
func NewProductService(usecase usecase.ProductUseCase) *ProductService {
	return &ProductService{
		userUseCase: usecase,
	}
}
