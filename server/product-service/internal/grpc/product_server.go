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

// func (s *ProductGRPCServer) GetProductSnapshot(ctx context.Context, req *productpb.GetProductRequest) (*productpb.ProductSnapshotResponse, error) {
// 	productID, err := uuid.Parse(req.GetProductId())
// 	if err != nil {
// 		return nil, err
// 	}

// 	product, err := s.Repo.FindByID(productID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Default: pakai harga global
// 	price := product.Price

// 	// Cek jika variantId diberikan
// 	if req.VariantId != "" {
// 		variantID, err := uuid.Parse(req.VariantId)
// 		if err == nil {
// 			variant, err := s.Repo.FindVariantByID(variantID)
// 			if err == nil && variant.IsActive {
// 				price = variant.Price // override harga dari variant
// 			}
// 		}
// 	}

// 	imageURL := ""
// 	if len(product.ProductImage) > 0 {
// 		imageURL = product.ProductImage[0].URL
// 	}

// 	return &productpb.ProductSnapshotResponse{
// 		Name:     product.Name,
// 		ImageUrl: imageURL,
// 		Price:    price,
// 	}, nil
// }
