package domain

import (
	"time"
)

type Order struct {
	ID          uint        `gorm:"primaryKey"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	Items       []OrderItem `json:"items"`
	Coupon      string      `json:"coupon"`
	TotalAmount float64     `json:"totalAmount"`
	CustomerID  int         `json:"-"`
	Status      string      `json:"status"`
}

type OrderItem struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ProductID int       `json:"-"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	Quantity  int       `json:"quantity"`
	Type      string    `json:"type"`
	OrderID   int       `json:"-"`
}
