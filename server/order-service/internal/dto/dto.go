package dto

type CheckoutRequest struct {
	CourierName  string  `json:"courierName" binding:"required"`
	ShippingCost float64 `json:"shippingCost" binding:"required"`
	Note         string  `json:"note"`
}

type ShippingCostRequest struct {
	DestinationID int    `json:"destinationId" binding:"required"`
	Weight        int    `json:"weight" binding:"required"`
	Courier       string `json:"courier" binding:"required"`
}

type CreateShipmentRequest struct {
	OrderID      string `json:"orderId" binding:"required"`
	TrackingCode string `json:"trackingCode" binding:"required"`
	Notes        string `json:"notes"`
}

type UpdateShipmentStatusRequest struct {
	Status      string `json:"status" binding:"required"`
	ShippedAt   string `json:"shippedAt"`
	DeliveredAt string `json:"deliveredAt"`
	Notes       string `json:"notes"`
}
