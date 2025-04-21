package repositories

import (
	"time"

	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetOrderByID(orderID uuid.UUID) (*models.Order, error)
	UpsertPayment(orderID uuid.UUID, status string, method string, paidAt *time.Time) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepo{db: db}
}

func (r *orderRepo) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepo) GetOrderByID(orderID uuid.UUID) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Items").Preload("Address").First(&order, "id = ?", orderID).Error
	return &order, err
}

func (r *orderRepo) UpsertPayment(orderID uuid.UUID, status string, method string, paidAt *time.Time) error {
	var payment models.Payment
	err := r.db.Where("order_id = ?", orderID).First(&payment).Error

	if err == gorm.ErrRecordNotFound {
		payment = models.Payment{
			ID:        uuid.New(),
			OrderID:   orderID,
			Method:    method,
			Status:    status,
			PaidAt:    paidAt,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		return r.db.Create(&payment).Error
	}

	payment.Status = status
	payment.Method = method
	payment.PaidAt = paidAt
	payment.UpdatedAt = time.Now()

	return r.db.Save(&payment).Error
}
