package services

import (
	"time"

	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/repositories"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/grpc"
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
	CreateOrder(userID, addressID uuid.UUID, items []models.OrderItem, note string, total float64, shippingCost float64, address models.Address) (*models.Order, error)
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
		return nil, "", errors.New("alamat utama tidak ditemukan, harap lengkapi alamat terlebih dahulu")
	}

	address := models.Address{ // snapshot
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
			items = append(items, models.OrderItem{...})
			total += item.Price * float64(item.Quantity)
			stockUpdates = append(stockUpdates, &productpb.StockUpdateItem{...})
		}
	}

	if len(items) == 0 {
		return nil, "", errors.New("No Item Selected")
	}
	if err := s.ProductGRPC.ReduceStock(stockUpdates); err != nil {
		return nil, "", err
	}

	order, err := s.CreateOrder(userID, items, total, req, address)
	if err != nil {
		return nil, "", err
	}

	_ = s.CartGRPC.ClearCart(userID.String())

	snapURL, err := s.GenerateSnapTransaction(order)
	return order, snapURL, err
}


func (s *OrderService) CreateOrder(
	userID uuid.UUID,
	items []models.OrderItem,
	total float64,
	req dto.CheckoutRequest,
	address models.Address,
) (*models.Order, error) {
	order := &models.Order{
		UserID:          userID,
		AddressID:       uuid.Parse(address.ID),
		Status:          "pending",
		TotalPrice:      total,
		ShippingCost:    Req.shippingCost,
		AmountToPay : total + Req.shippingCost,
		Note:            Req.note,
		Items:           items,
		CourierName : req.CourierName,
		ShippingName:    address.Name,
		ShippingAddress: address.Address,
		City:            address.City,
		Province:        address.Province,
		Zipcode:         address.Zipcode,
		Phone:           address.Phone,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	return order, s.Repo.CreateOrder(order)
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
