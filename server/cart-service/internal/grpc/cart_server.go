package grpc

import (
	"context"

	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/services"
	cartpb "github.com/fiqrioemry/microservice-ecommerce/server/proto/cart"
	"github.com/google/uuid"
)

type CartGRPCServer struct {
	cartpb.UnimplementedCartServiceServer
	Service services.CartService
}

func NewCartGRPCServer(service services.CartService) *CartGRPCServer {
	return &CartGRPCServer{Service: service}
}

func (s *CartGRPCServer) ClearCart(ctx context.Context, req *cartpb.ClearCartRequest) (*cartpb.EmptyCartResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}

	if err := s.Service.ClearUserCart(userID.String()); err != nil {
		return nil, err
	}

	return &cartpb.EmptyCartResponse{Message: "Cart cleared"}, nil
}

func (s *CartGRPCServer) GetCartForCheckout(ctx context.Context, req *cartpb.GetCartRequest) (*cartpb.CartResponse, error) {
	userID := req.GetUserId()

	cart, err := s.Service.GetUserCart(userID)
	if err != nil {
		return nil, err
	}

	var items []*cartpb.CartItem
	for _, item := range cart.Items {
		variantID := ""
		if item.VariantID != nil {
			variantID = item.VariantID.String()
		}

		items = append(items, &cartpb.CartItem{
			ProductId:   item.ProductID.String(),
			VariantId:   variantID,
			ProductName: item.ProductName,
			ImageUrl:    item.ImageURL,
			Price:       item.Price,
			Quantity:    int32(item.Quantity),
			IsChecked:   item.IsChecked,
		})
	}

	return &cartpb.CartResponse{Items: items}, nil
}
