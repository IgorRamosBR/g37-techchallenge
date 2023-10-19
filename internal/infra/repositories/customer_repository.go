package repositories

import (
	"fmt"
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/ports"
	"g37-lanchonete/internal/infra/clients"
)

type customerRepository struct {
	client clients.SQLClient
}

func NewcustomerRepository(client clients.SQLClient) ports.CustomerRepository {
	return customerRepository{
		client: client,
	}
}

func (r customerRepository) SaveCustomer(customer models.Customer) error {
	err := r.client.Save(&customer)
	if err != nil {
		return fmt.Errorf("failed to save customer, error %v", err)
	}

	return nil
}
func (r customerRepository) FindCustomerByCPF(cpf string) (models.Customer, error) {

	return models.Customer{
		Name:  "Joao",
		Cpf:   "11133322200",
		Email: "joao@gmail.com",
	}, nil
}
