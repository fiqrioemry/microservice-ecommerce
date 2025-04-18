package services

import (
	"errors"

	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/repositories"
	"github.com/google/uuid"
)

type CartService interface {
	GetUserCart(userID string) (*models.Cart, error)
	AddToCart(userID string, req dto.AddToCartRequest, productSnapshot dto.ProductSnapshot) error
	UpdateCartItem(itemID uuid.UUID, req dto.UpdateCartItemRequest) error
	RemoveCartItem(itemID uuid.UUID) error
	ClearUserCart(userID string) error
}

type cartService struct {
	repo repositories.CartRepository
}

func NewCartService(repo repositories.CartRepository) CartService {
	return &cartService{repo}
}

func (s *cartService) GetUserCart(userID string) (*models.Cart, error) {
	return s.repo.GetCartWithItems(userID)
}

func (s *cartService) AddToCart(userID string, req dto.AddToCartRequest, product dto.ProductSnapshot) error {
	cart, err := s.repo.GetOrCreateCart(userID)
	if err != nil {
		return err
	}

	item := &models.CartItem{
		ID:          uuid.New(),
		CartID:      cart.ID,
		ProductID:   req.ProductID,
		VariantID:   req.VariantID,
		ProductName: product.Name,
		ImageURL:    product.ImageURL,
		Price:       product.Price,
		Quantity:    req.Quantity,
		IsChecked:   true,
	}
	return s.repo.AddItem(item)
}

func (s *cartService) UpdateCartItem(itemID uuid.UUID, req dto.UpdateCartItemRequest) error {
	item, err := s.repo.FindItemByID(itemID)
	if err != nil {
		return errors.New("cart item not found")
	}
	item.Quantity = req.Quantity
	item.IsChecked = req.IsChecked
	return s.repo.UpdateItem(item)
}

func (s *cartService) RemoveCartItem(itemID uuid.UUID) error {
	return s.repo.DeleteItem(itemID)
}

func (s *cartService) ClearUserCart(userID string) error {
	cart, err := s.repo.GetOrCreateCart(userID)
	if err != nil {
		return err
	}
	return s.repo.ClearCart(cart.ID)
}
