package ports

import (
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/services/dto"
)

type CustomerService interface {
	CreateCustomer(customer dto.CustomerDTO) error
	GetCustomerByCPF(cpf string) (models.Customer, error)
}
