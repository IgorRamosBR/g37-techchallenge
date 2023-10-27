package domain

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name   string `gorm:"size:100"`
	Cpf    string `gorm:"size:11"`
	Email  string `gorm:"size:100"`
	Orders []Order
}
