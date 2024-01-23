package entities

import (
	"time"
)

type Order struct {
	ID          int         `json:"id"`
	Items       []OrderItem `json:"items"`
	Coupon      string      `json:"coupon"`
	TotalAmount float64     `json:"totalAmount"`
	Customer    Customer    `json:"customer"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"createdAt"`
}

type OrderItem struct {
	ID       int     `json:"id"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
	Type     string  `json:"type"`
}
