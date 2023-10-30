package domain

import "time"

type Product struct {
	ID          int       `gorm:"primaryKey"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name" gorm:"size:100"`
	SkuId       string    `json:"skuId" gorm:"size:50"`
	Description string    `json:"description" gorm:"size:2000"`
	Category    string    `json:"category" gorm:"size:60"`
	Price       float64   `json:"price"`
}
