package dto

import (
	"time"

	"github.com/asaskevich/govalidator"
)

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
	QrData       string `json:"qr_data"`
	StoreOrderId string `json:"in_store_order_id"`
}

type PaymentQRCode struct {
	QRCode string `json:"qrcode"`
}

type PaymentNotificationDTO struct {
	Id          string      `json:"id"`
	LiveMode    bool        `json:"liveMode"`
	Type        string      `json:"type" valid:"in(payment),required~Type is invalid"`
	DateCreated time.Time   `json:"dateCreated"`
	UserId      int         `json:"userId"`
	ApiVersion  string      `json:"apiVersion"`
	Action      string      `json:"action"`
	Data        PaymentData `json:"data"`
}

type PaymentData struct {
	Id string `json:"id" valid:"required,numeric"`
}

func (p PaymentNotificationDTO) ValidatePaymentNotification() (bool, error) {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return false, err
	}

	return true, nil
}
