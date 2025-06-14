package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetOrCreateCart(userID string) (*models.Cart, error)
	GetCartWithItems(userID string) (*models.Cart, error)
	AddItem(item *models.CartItem) error
	UpdateItem(item *models.CartItem) error
	DeleteItem(itemID uuid.UUID) error
	ClearCart(cartID uuid.UUID) error
	PreloadCartItems(cart *models.Cart) error
	FindItemByID(itemID uuid.UUID) (*models.CartItem, error)
	FindItemByCartProductVariant(cartID, productID uuid.UUID, variantID *uuid.UUID) (*models.CartItem, error)
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) GetOrCreateCart(userID string) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Where("user_id = ?", userID).First(&cart).Error
	if err == gorm.ErrRecordNotFound {
		cart = models.Cart{
			ID:     uuid.New(),
			UserID: uuid.MustParse(userID),
		}
		if err := r.db.Create(&cart).Error; err != nil {
			return nil, err
		}
		return &cart, nil
	}
	return &cart, err
}

func (r *cartRepository) FindItemByCartProductVariant(
	cartID, productID uuid.UUID, variantID *uuid.UUID,
) (*models.CartItem, error) {
	var item models.CartItem
	query := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID)

	if variantID != nil {
		query = query.Where("variant_id = ?", *variantID)
	} else {
		query = query.Where("variant_id IS NULL")
	}

	err := query.First(&item).Error

	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *cartRepository) GetCartWithItems(userID string) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Items").Where("user_id = ?", userID).First(&cart).Error
	return &cart, err
}

func (r *cartRepository) AddItem(item *models.CartItem) error {
	return r.db.Create(item).Error
}

func (r *cartRepository) UpdateItem(item *models.CartItem) error {
	return r.db.Save(item).Error
}

func (r *cartRepository) DeleteItem(itemID uuid.UUID) error {
	return r.db.Delete(&models.CartItem{}, "id = ?", itemID).Error
}

func (r *cartRepository) ClearCart(cartID uuid.UUID) error {
	return r.db.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error
}

func (r *cartRepository) FindItemByID(itemID uuid.UUID) (*models.CartItem, error) {
	var item models.CartItem
	err := r.db.First(&item, "id = ?", itemID).Error
	return &item, err
}

func (r *cartRepository) PreloadCartItems(cart *models.Cart) error {
	return r.db.Preload("Items").First(cart, "id = ?", cart.ID).Error
}
