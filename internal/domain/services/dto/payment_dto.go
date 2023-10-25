package dto

type PaymentQRCodeRequest struct {
	ExternalReference string               `json:"external_reference"`
	Title             string               `json:"title"`
	NotificationURL   string               `json:"notification_url"`
	TotalAmount       float64              `json:"total_amount"`
	Items             []PaymentItemRequest `json:"items"`
	Sponsor           string               `json:"sponsor"`
}

type PaymentItemRequest struct {
	SkuNumber   string  `json:"sku_number"`
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	UnitMeasure string  `json:"unit_measure"`
	TotalAmount float64 `json:"total_amount"`
}

type SponsorRequest struct {
	Id string `json:"id"`
}

type PaymentQRCodeResponse struct {
	QrData       string `json:"qrData"`
	StoreOrderId string `json:"storeOrderId"`
}

type PaymentQRCode struct {
	QRCode string `json:"qrcode"`
}
