package grpc

import (
	"context"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
	productpb "github.com/fiqrioemry/microservice-ecommerce/server/proto/product"
	"github.com/google/uuid"
)

type ProductGRPCServer struct {
	productpb.UnimplementedProductServiceServer
	Repo repositories.ProductRepository
}

func NewProductGRPCServer(repo repositories.ProductRepository) *ProductGRPCServer {
	return &ProductGRPCServer{Repo: repo}
}

func (s *ProductGRPCServer) GetProductSnapshot(ctx context.Context, req *productpb.GetProductRequest) (*productpb.ProductSnapshotResponse, error) {
	productID, err := uuid.Parse(req.GetProductId())
	if err != nil {
		return nil, err
	}

	product, err := s.Repo.FindByID(productID)
	if err != nil {
		return nil, err
	}

	imageURL := ""
	if len(product.ProductImage) > 0 {
		imageURL = product.ProductImage[0].URL
	}

	return &productpb.ProductSnapshotResponse{
		Name:     product.Name,
		ImageUrl: imageURL,
		Price:    product.Price,
	}, nil
}
