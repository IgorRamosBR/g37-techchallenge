package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"size:100"`
	SkuId       string `gorm:"size:50"`
	Description string `gorm:"size:2000"`
	Category    string `gorm:"size:60"`
	Price       float64
}
