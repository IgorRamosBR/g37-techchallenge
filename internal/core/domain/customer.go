package domain

import "time"

type Customer struct {
	ID        int       `gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name" gorm:"size:100"`
	Cpf       string    `json:"cpf" gorm:"size:11"`
	Email     string    `json:"email" gorm:"size:100"`
	Orders    []Order   `json:"orders,omitempty"`
}
