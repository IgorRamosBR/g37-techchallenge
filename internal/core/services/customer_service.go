package services

import (
	"g37-lanchonete/internal/core/entities"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"g37-lanchonete/internal/infra/gateways"
	"time"

	log "github.com/sirupsen/logrus"
)

type customerService struct {
	customerRepositoryGateway gateways.CustomerRepositoryGateway
}

func NewCustomerService(customerRepository gateways.CustomerRepositoryGateway) ports.CustomerService {
	return customerService{
		customerRepositoryGateway: customerRepository,
	}
}

func (s customerService) CreateCustomer(customerDTO dto.CustomerDTO) error {
	customer := customerDTO.ToCustomer()
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	err := s.customerRepositoryGateway.SaveCustomer(customer)
	if err != nil {
		log.Errorf("failed to save customer, error: %v", err)
		return err
	}

	return nil
}

func (s customerService) GetCustomerById(id int) (entities.Customer, error) {
	customer, err := s.customerRepositoryGateway.FindCustomerById(id)
	if err != nil {
		log.Errorf("failed to get customer by id [%d], error: %v", id, err)
		return entities.Customer{}, err
	}

	return customer, nil
}

func (s customerService) GetCustomerByCPF(cpf string) (entities.Customer, error) {
	customer, err := s.customerRepositoryGateway.FindCustomerByCPF(cpf)
	if err != nil {
		log.Errorf("failed to get customer by cpf [%s], error: %v", cpf, err)
		return entities.Customer{}, err
	}

	return customer, nil
}
