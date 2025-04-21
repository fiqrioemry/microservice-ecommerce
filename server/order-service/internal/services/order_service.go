package services

import (
	"errors"
	"time"

	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/repositories"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/grpc"
	productpb "github.com/fiqrioemry/microservice-ecommerce/server/proto/product"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type OrderServiceInterface interface {
	GetCart(userID uuid.UUID) (*models.Cart, error) // komunikasi antar service via gRPC
	GetAllOrders() ([]models.Order, error)
	GetUserOrdersByID(userID uuid.UUID) ([]models.Order, error)
	//
	GenerateSnapTransaction(order *models.Order) (string, error)
	UpdatePaymentStatus(orderID string, transactionStatus string, paymentType string) error
	CreateOrderWithMainAddress(userID uuid.UUID, cart *models.Cart, req dto.CheckoutRequest) (*models.Order, string, error)
	CreateOrder(userID, addressID uuid.UUID, items []models.OrderItem, req dto.CheckoutRequest, total float64, address models.Address) (*models.Order, error)
	CreateShipment(req dto.CreateShipmentRequest) (*models.Shipment, error)
	GetShipmentByOrderID(orderID uuid.UUID) (*models.Shipment, error)
	UpdateShipmentStatus(orderID uuid.UUID, req dto.UpdateShipmentStatusRequest) error
}

type OrderService struct {
	Repo        repositories.OrderRepository
	CartGRPC    *grpc.CartGRPCClient
	UserGRPC    *grpc.UserGRPCClient
	ProductGRPC *grpc.ProductGRPCClient
}

func NewOrderService(repo repositories.OrderRepository, cartClient *grpc.CartGRPCClient, userClient *grpc.UserGRPCClient, productClient *grpc.ProductGRPCClient) OrderServiceInterface {
	return &OrderService{
		Repo: repo, CartGRPC: cartClient, UserGRPC: userClient, ProductGRPC: productClient,
	}
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

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.Repo.GetAllUserOrders()
}

func (s *OrderService) GetUserOrdersByID(userID uuid.UUID) ([]models.Order, error) {
	return s.Repo.GetUserOrdersByID(userID)
}

func (s *OrderService) CreateOrderWithMainAddress(userID uuid.UUID, cart *models.Cart, req dto.CheckoutRequest) (*models.Order, string, error) {
	resp, err := s.UserGRPC.GetMainAddress(userID.String())
	if err != nil {
		return nil, "", errors.New("address not found, please select an address")
	}

	addressID, _ := uuid.Parse(resp.Id)
	address := models.Address{
		Name:     resp.Name,
		Address:  resp.Address,
		City:     resp.City,
		Province: resp.Province,
		Zipcode:  resp.Zipcode,
		Phone:    resp.Phone,
	}

	var items []models.OrderItem
	var total float64
	var stockUpdates []*productpb.StockUpdateItem

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

			stockUpdates = append(stockUpdates, &productpb.StockUpdateItem{
				ProductId: item.ProductID.String(),
				VariantId: func() string {
					if item.VariantID != nil {
						return item.VariantID.String()
					}
					return ""
				}(),
				Quantity: int32(item.Quantity),
			})
		}
	}

	if len(items) == 0 {
		return nil, "", errors.New("no item selected")
	}

	if err := s.ProductGRPC.ReduceStock(stockUpdates); err != nil {
		return nil, "", err
	}

	order, err := s.CreateOrder(userID, addressID, items, req, total, address)
	if err != nil {
		return nil, "", err
	}

	_ = s.CartGRPC.ClearCart(userID.String())

	snapURL, err := s.GenerateSnapTransaction(order)
	return order, snapURL, err
}

func (s *OrderService) CreateOrder(userID, addressID uuid.UUID, items []models.OrderItem, req dto.CheckoutRequest, total float64, address models.Address) (*models.Order, error) {
	order := &models.Order{
		UserID:          userID,
		AddressID:       addressID,
		Status:          "pending",
		TotalPrice:      total,
		ShippingCost:    req.ShippingCost,
		AmountToPay:     total + req.ShippingCost,
		Note:            req.Note,
		Items:           items,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		CourierName:     req.CourierName,
		ShippingName:    address.Name,
		ShippingAddress: address.Address,
		City:            address.City,
		Province:        address.Province,
		Zipcode:         address.Zipcode,
		Phone:           address.Phone,
	}
	return order, s.Repo.CreateOrder(order)
}

func (s *OrderService) GenerateSnapTransaction(order *models.Order) (string, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  order.ID.String(),
			GrossAmt: int64(order.AmountToPay),
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

func (s *OrderService) CreateShipment(req dto.CreateShipmentRequest) (*models.Shipment, error) {
	orderID, err := uuid.Parse(req.OrderID)
	if err != nil {
		return nil, errors.New("invalid order ID")
	}
	shipment := &models.Shipment{
		OrderID:      orderID,
		TrackingCode: req.TrackingCode,
		Status:       "pending",
		Notes:        req.Notes,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := s.Repo.CreateShipment(shipment); err != nil {
		return nil, err
	}
	return shipment, nil
}

func (s *OrderService) GetShipmentByOrderID(orderID uuid.UUID) (*models.Shipment, error) {
	return s.Repo.GetShipmentByOrderID(orderID)
}

func (s *OrderService) UpdateShipmentStatus(orderID uuid.UUID, req dto.UpdateShipmentStatusRequest) error {
	update := map[string]interface{}{
		"status":     req.Status,
		"notes":      req.Notes,
		"updated_at": time.Now(),
	}
	if req.ShippedAt != "" {
		if t, err := time.Parse(time.RFC3339, req.ShippedAt); err == nil {
			update["shipped_at"] = &t
		}
	}
	if req.DeliveredAt != "" {
		if t, err := time.Parse(time.RFC3339, req.DeliveredAt); err == nil {
			update["delivered_at"] = &t
		}
	}
	return s.Repo.UpdateShipmentStatus(orderID, update)
}
