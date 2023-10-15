package ports

import "g37-lanchonete/internal/domain/models"

type CustomerService interface {
	CreateCustomer(customer models.Customer) error
	GetCustomerByCPF(cpf string) (models.Customer, error)
}
