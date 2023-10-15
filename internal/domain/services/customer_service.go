package services

import (
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/ports"
)

type customerService struct {
	customerRepository ports.CustomerRepository
}

func NewCustomerService(customerRepository ports.CustomerRepository) ports.CustomerService {
	return customerService{
		customerRepository: customerRepository,
	}
}

func (s customerService) CreateCustomer(customer models.Customer) error {
	return nil
}

func (s customerService) GetCustomerByCPF(cpf string) (models.Customer, error) {
	customer, err := s.customerRepository.FindCustomerByCPF(cpf)
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}
