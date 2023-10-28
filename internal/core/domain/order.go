package domain

import (
	"time"
)

type Order struct {
	ID          uint      `gorm:"primaryKey"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Items       []OrderItem
	Coupon      string
	Discount    float64
	TotalAmount float64
	CustomerID  uint
	Status      string
}

type OrderItem struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
	Quantity  int
	Type      string
	OrderID   uint
}
