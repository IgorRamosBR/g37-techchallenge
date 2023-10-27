package repositories

import (
	"fmt"
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/infra/clients"
)

type customerRepository struct {
	client clients.SQLClient
}

func NewCustomerRepository(client clients.SQLClient) ports.CustomerRepository {
	return customerRepository{
		client: client,
	}
}

func (r customerRepository) FindCustomerById(id string) (domain.Customer, error) {
	var customer domain.Customer

	err := r.client.FindById(id, customer)
	if err != nil {
		return domain.Customer{}, fmt.Errorf("failed to find customer by id [%s], error %v", id, err)
	}

	return customer, nil
}

func (r customerRepository) FindCustomerByCPF(cpf string) (domain.Customer, error) {
	customer := domain.Customer{Cpf: cpf}

	err := r.client.FindFirst(&customer, "cpf = ?", cpf)
	if err != nil {
		return domain.Customer{}, fmt.Errorf("failed to find customer by cpf [%s], error %v", cpf, err)
	}

	return customer, nil
}

func (r customerRepository) SaveCustomer(customer domain.Customer) error {
	err := r.client.Save(&customer)
	if err != nil {
		return fmt.Errorf("failed to save customer, error %v", err)
	}

	return nil
}
