package services

import (
	"errors"

	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartService interface {
	GetUserCart(userID string) (*models.Cart, error)
	AddToCart(userID string, req dto.AddToCartRequest, productSnapshot dto.ProductSnapshot) error

	UpdateCartItem(itemID uuid.UUID, req dto.UpdateCartItemRequest, productSnapshot dto.ProductSnapshot) error
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
	cart, err := s.repo.GetOrCreateCart(userID)
	if err != nil {
		return nil, err
	}

	err = s.repo.PreloadCartItems(cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (s *cartService) AddToCart(userID string, req dto.AddToCartRequest, snapshot dto.ProductSnapshot) error {
	cart, err := s.repo.GetOrCreateCart(userID)
	if err != nil {
		return err
	}

	existing, err := s.repo.FindItemByCartProductVariant(cart.ID, req.ProductID, req.VariantID)
	if err == nil {
		totalQty := existing.Quantity + req.Quantity
		if totalQty > snapshot.Stock {
			return errors.New("quantity exceeds available stock")
		}
		existing.Quantity = totalQty
		return s.repo.UpdateItem(existing)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if req.Quantity > snapshot.Stock {
			return errors.New("quantity exceeds available stock")
		}
		newItem := &models.CartItem{
			ID:          uuid.New(),
			CartID:      cart.ID,
			ProductID:   req.ProductID,
			VariantID:   req.VariantID,
			ProductName: snapshot.Name,
			ImageURL:    snapshot.ImageURL,
			Price:       snapshot.Price,
			Quantity:    req.Quantity,
			IsChecked:   true,
		}
		return s.repo.AddItem(newItem)
	}

	return err
}

func (s *cartService) UpdateCartItem(itemID uuid.UUID, req dto.UpdateCartItemRequest, snapshot dto.ProductSnapshot) error {
	item, err := s.repo.FindItemByID(itemID)
	if err != nil {
		return errors.New("cart item not found")
	}

	if req.Quantity > snapshot.Stock {
		return errors.New("quantity exceeds available stock")
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
