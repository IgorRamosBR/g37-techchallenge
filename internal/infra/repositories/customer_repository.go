package repositories

import (
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/ports"
)

type customerRepository struct {
}

func NewcustomerRepository() ports.CustomerRepository {
	return customerRepository{}
}

func (r customerRepository) SaveCustomer(customer models.Customer) error {

	return nil
}
func (r customerRepository) FindCustomerByCPF(cpf string) (models.Customer, error) {

	return models.Customer{
		Name:  "Joao",
		Cpf:   "11133322200",
		Email: "joao@gmail.com",
	}, nil
}
