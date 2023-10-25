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

func (r customerRepository) FindCustomerById(id string) (models.Customer, error) {
	var customer models.Customer

	err := r.client.FindById(id, customer)
	if err != nil {
		return models.Customer{}, fmt.Errorf("failed to find customer by id [%s], error %v", id, err)
	}

	return customer, nil
}

func (r customerRepository) FindCustomerByCPF(cpf string) (models.Customer, error) {
	customer := models.Customer{Cpf: cpf}

	err := r.client.FindFirst(&customer, "cpf = ?", cpf)
	if err != nil {
		return models.Customer{}, fmt.Errorf("failed to find customer by cpf [%s], error %v", cpf, err)
	}

	return customer, nil
}

func (r customerRepository) SaveCustomer(customer models.Customer) error {
	err := r.client.Save(&customer)
	if err != nil {
		return fmt.Errorf("failed to save customer, error %v", err)
	}

	return nil
}
