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
