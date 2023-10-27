package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Items       []OrderItem
	Coupon      string
	Discount    float64
	TotalAmount float64
	CustomerID  uint
	Status      string
}

type OrderItem struct {
	gorm.Model
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
	Quantity  int
	Type      string
	OrderID   uint
}
