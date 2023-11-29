package dto

type OrderCreationResponse struct {
	QRCode  string `json:"qrCode"`
	OrderID uint   `json:"orderId"`
}
