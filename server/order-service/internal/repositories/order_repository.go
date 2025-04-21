package repositories

import (
	"time"

	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetAllUserOrders() ([]models.Order, error)
	GetUserOrdersByID(userID uuid.UUID) ([]models.Order, error)
	UpsertPayment(orderID uuid.UUID, status string, method string, paidAt *time.Time) error
	CreateShipment(shipment *models.Shipment) error
	GetShipmentByOrderID(orderID uuid.UUID) (*models.Shipment, error)
	UpdateShipmentStatus(orderID uuid.UUID, update map[string]interface{}) error
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

func (r *orderRepo) GetAllUserOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Items").Preload("Address").Find(&orders).Error
	return orders, err
}
func (r *orderRepo) GetUserOrdersByID(userID uuid.UUID) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Items").Preload("Address").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
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

func (r *orderRepo) CreateShipment(shipment *models.Shipment) error {
	return r.db.Create(shipment).Error
}

func (r *orderRepo) GetShipmentByOrderID(orderID uuid.UUID) (*models.Shipment, error) {
	var shipment models.Shipment
	err := r.db.First(&shipment, "order_id = ?", orderID).Error
	return &shipment, err
}

func (r *orderRepo) UpdateShipmentStatus(orderID uuid.UUID, update map[string]interface{}) error {
	return r.db.Model(&models.Shipment{}).Where("order_id = ?", orderID).Updates(update).Error
}
