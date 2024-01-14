package domain

import "time"

type Customer struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Email     string    `json:"email"`
	Orders    []Order   `json:"orders,omitempty"`
}
