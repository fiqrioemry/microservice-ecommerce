package services

import (
	"errors"
	"time"

	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/grpc"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/repositories"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type OrderServiceInterface interface {
	GetCart(userID uuid.UUID) (*models.Cart, error)
	GetOrderDetail(orderID uuid.UUID) (*models.Order, error)
	GenerateSnapTransaction(order *models.Order) (string, error)
	UpdatePaymentStatus(orderID string, transactionStatus string, paymentType string) error
	CreateOrder(userID, addressID uuid.UUID, items []models.OrderItem, note string, total float64, shippingCost float64, address models.Address) (*models.Order, error)
	CreateOrderFromCart(userID, addressID uuid.UUID, cart *models.Cart, note string, shippingCost float64) (*models.Order, error)
}

type OrderService struct {
	Repo     repositories.OrderRepository
	CartGRPC *grpc.CartGRPCClient
	UserGRPC *grpc.UserGRPCClient
}

func NewOrderService(repo repositories.OrderRepository, cartClient *grpc.CartGRPCClient, userClient *grpc.UserGRPCClient) OrderServiceInterface {
	return &OrderService{Repo: repo, CartGRPC: cartClient, UserGRPC: userClient}
}

func (s *OrderService) CreateOrderFromCart(
	userID, addressID uuid.UUID,
	cart *models.Cart,
	note string,
	shippingCost float64,
) (*models.Order, error) {
	var items []models.OrderItem
	var total float64

	for _, item := range cart.Items {
		if item.IsChecked {
			items = append(items, models.OrderItem{
				ProductID:   item.ProductID,
				VariantID:   item.VariantID,
				ProductName: item.ProductName,
				ImageURL:    item.ImageURL,
				Price:       item.Price,
				Quantity:    item.Quantity,
			})
			total += item.Price * float64(item.Quantity)
		}
	}

	if len(items) == 0 {
		return nil, errors.New("no items selected for checkout")
	}

	resp, err := s.UserGRPC.GetAddressByID(addressID.String())
	if err != nil {
		return nil, err
	}

	snapshot := models.Address{
		Name:     resp.Name,
		Address:  resp.Address,
		City:     resp.City,
		Province: resp.Province,
		Zipcode:  resp.Zipcode,
		Phone:    resp.Phone,
	}

	return s.CreateOrder(userID, addressID, items, note, total, shippingCost, snapshot)
}

func (s *OrderService) CreateOrder(
	userID, addressID uuid.UUID,
	items []models.OrderItem,
	note string,
	total float64,
	shippingCost float64,
	address models.Address,
) (*models.Order, error) {
	order := &models.Order{
		UserID:          userID,
		AddressID:       addressID,
		Status:          "pending",
		TotalPrice:      total,
		ShippingCost:    shippingCost,
		Note:            note,
		Items:           items,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		ShippingName:    address.Name,
		ShippingAddress: address.Address,
		City:            address.City,
		Province:        address.Province,
		Zipcode:         address.Zipcode,
		Phone:           address.Phone,
	}
	return order, s.Repo.CreateOrder(order)
}
func (s *OrderService) GetOrderDetail(orderID uuid.UUID) (*models.Order, error) {
	return s.Repo.GetOrderByID(orderID)
}

func (s *OrderService) GetCart(userID uuid.UUID) (*models.Cart, error) {
	cartItems, err := s.CartGRPC.GetCart(userID.String())
	if err != nil {
		return nil, err
	}

	var items []models.CartItem
	for _, item := range cartItems {
		productID, _ := uuid.Parse(item.ProductId)
		variantID := uuid.Nil
		if item.VariantId != "" {
			variantID, _ = uuid.Parse(item.VariantId)
		}

		items = append(items, models.CartItem{
			ProductID:   productID,
			VariantID:   &variantID,
			ProductName: item.ProductName,
			ImageURL:    item.ImageUrl,
			Price:       item.Price,
			Quantity:    int(item.Quantity),
			IsChecked:   item.IsChecked,
		})
	}

	return &models.Cart{Items: items}, nil
}

func (s *OrderService) GenerateSnapTransaction(order *models.Order) (string, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  order.ID.String(),
			GrossAmt: int64(order.TotalPrice),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: order.ShippingName,
			Phone: order.Phone,
		},
	}

	snapResp, err := config.SnapClient.CreateTransaction(req)
	if err != nil {
		return "", err
	}
	return snapResp.RedirectURL, nil
}

func (s *OrderService) UpdatePaymentStatus(orderID string, transactionStatus string, paymentType string) error {
	id, err := uuid.Parse(orderID)
	if err != nil {
		return err
	}

	var paidAt *time.Time
	status := "pending"

	switch transactionStatus {
	case "settlement", "capture":
		now := time.Now()
		paidAt = &now
		status = "paid"
	case "expire", "cancel":
		status = "failed"
	}

	return s.Repo.UpsertPayment(id, status, paymentType, paidAt)
}
