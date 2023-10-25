package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Items       []OrderItem
	Coupon      string
	Discount    float64
	TotalAmount float64
	Customer    Customer
	Status      string
}

type OrderItem struct {
	gorm.Model
	Product  Product
	Quantity int
	Type     string
}
