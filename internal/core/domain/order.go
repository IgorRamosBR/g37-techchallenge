package domain

import (
	"time"
)

type Order struct {
	ID          int         `json:"id"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	Items       []OrderItem `json:"items"`
	Coupon      string      `json:"coupon"`
	TotalAmount float64     `json:"totalAmount"`
	Customer    Customer    `json:"customer"`
	Status      string      `json:"status"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	Product   Product   `json:"product"`
	Quantity  int       `json:"quantity"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
