package grpc

import (
	"context"
	"fmt"

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
	variantIDStr := req.GetVariantId()

	if variantIDStr != "" {
		variantID, err := uuid.Parse(variantIDStr)
		if err != nil {
			return nil, err
		}
		variant, err := s.Repo.FindVariantByID(variantID)
		if err != nil {
			return nil, err
		}
		imageURL := variant.ImageURL
		if imageURL == "" && len(variant.Product.ProductImage) > 0 {
			imageURL = variant.Product.ProductImage[0].URL
		}

		return &productpb.ProductSnapshotResponse{
			Name:     variant.Product.Name,
			ImageUrl: imageURL,
			Price:    variant.Price,
			Stock:    int32(variant.Stock),
		}, nil

	}

	// fallback jika tidak ada variant
	productID, err := uuid.Parse(req.GetProductId())
	if err != nil {
		return nil, err
	}
	product, err := s.Repo.FindByID(productID)
	if err != nil {
		return nil, err
	}

	price := 0.0
	stock := 0
	if len(product.ProductVariant) > 0 {
		price = product.ProductVariant[0].Price
		stock = product.ProductVariant[0].Stock
	}

	imageURL := ""
	if len(product.ProductImage) > 0 {
		imageURL = product.ProductImage[0].URL
	}

	return &productpb.ProductSnapshotResponse{
		Name:     product.Name,
		ImageUrl: imageURL,
		Price:    price,
		Stock:    int32(stock),
	}, nil
}

func (s *ProductGRPCServer) GetMultipleProductSnapshots(ctx context.Context, req *productpb.GetMultipleProductRequest) (*productpb.MultipleProductSnapshotResponse, error) {
	var snapshots []*productpb.ProductSnapshot
	for _, item := range req.GetItems() {
		id, err := uuid.Parse(item.GetProductId())
		if err != nil {
			continue
		}
		product, err := s.Repo.FindByID(id)
		if err != nil {
			continue
		}

		price := 0.0
		stock := int32(0)
		imageURL := ""
		if len(product.ProductVariant) > 0 {
			price = product.ProductVariant[0].Price
			stock = int32(product.ProductVariant[0].Stock)
			imageURL = product.ProductVariant[0].ImageURL
		} else if len(product.ProductImage) > 0 {
			imageURL = product.ProductImage[0].URL
		}

		snapshots = append(snapshots, &productpb.ProductSnapshot{
			Id:       product.ID.String(),
			Name:     product.Name,
			ImageUrl: imageURL,
			Price:    price,
			Stock:    stock,
		})
	}
	return &productpb.MultipleProductSnapshotResponse{
		Products: snapshots,
	}, nil
}

func (s *ProductGRPCServer) CheckProductAvailability(ctx context.Context, req *productpb.CheckAvailabilityRequest) (*productpb.CheckAvailabilityResponse, error) {
	productID, err := uuid.Parse(req.GetProductId())
	if err != nil {
		return nil, err
	}

	product, err := s.Repo.FindByID(productID)
	if err != nil {
		return nil, err
	}

	inStock := false
	for _, variant := range product.ProductVariant {
		if variant.Stock > 0 {
			inStock = true
			break
		}
	}

	return &productpb.CheckAvailabilityResponse{
		IsActive: product.IsActive,
		InStock:  inStock,
	}, nil
}

func (s *ProductGRPCServer) UpdateProductStock(ctx context.Context, req *productpb.UpdateStockRequest) (*productpb.EmptyResponse, error) {
	for _, item := range req.Items {
		productID, err := uuid.Parse(item.ProductId)
		if err != nil {
			return nil, err
		}

		variantID := uuid.Nil
		if item.VariantId != "" {
			variantID, err = uuid.Parse(item.VariantId)
			if err != nil {
				return nil, err
			}
		}

		// Ambil variant (kalau ada)
		if variantID != uuid.Nil {
			variant, err := s.Repo.FindVariantByID(variantID)
			if err != nil {
				return nil, err
			}
			if variant.Stock < int(item.Quantity) {
				return nil, fmt.Errorf("insufficient stock for variant %s", variantID)
			}
			variant.Stock -= int(item.Quantity)
			if err := s.Repo.UpdateVariant(variant); err != nil {
				return nil, err
			}
		} else {
			// fallback ke variant pertama dari produk
			product, err := s.Repo.FindByID(productID)
			if err != nil {
				return nil, err
			}
			if len(product.ProductVariant) == 0 {
				return nil, fmt.Errorf("no variant found for product %s", productID)
			}
			variant := &product.ProductVariant[0]
			if variant.Stock < int(item.Quantity) {
				return nil, fmt.Errorf("insufficient stock for product %s", productID)
			}
			variant.Stock -= int(item.Quantity)
			if err := s.Repo.UpdateVariant(variant); err != nil {
				return nil, err
			}
		}
	}

	return &productpb.EmptyResponse{}, nil
}
