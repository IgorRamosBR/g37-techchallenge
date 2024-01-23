package dto

type OrderCreationResponse struct {
	QRCode  string `json:"qrCode"`
	OrderID int    `json:"orderId"`
}
