package services

import (
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/ports"
	"g37-lanchonete/internal/domain/services/dto"

	log "github.com/sirupsen/logrus"
)

type customerService struct {
	customerRepository ports.CustomerRepository
}

func NewCustomerService(customerRepository ports.CustomerRepository) ports.CustomerService {
	return customerService{
		customerRepository: customerRepository,
	}
}

func (s customerService) CreateCustomer(customerDTO dto.CustomerDTO) error {
	customer := customerDTO.ToCustomer()

	err := s.customerRepository.SaveCustomer(customer)
	if err != nil {
		log.Errorf("failed to save customer, error: %v", err)
		return err
	}

	return nil
}

func (s customerService) GetCustomerById(id string) (models.Customer, error) {
	customer, err := s.customerRepository.FindCustomerById(id)
	if err != nil {
		log.Errorf("failed to get customer by id [%s], error: %v", id, err)
		return models.Customer{}, err
	}

	return customer, nil
}

func (s customerService) GetCustomerByCPF(cpf string) (models.Customer, error) {
	customer, err := s.customerRepository.FindCustomerByCPF(cpf)
	if err != nil {
		log.Errorf("failed to get customer by cpf [%s], error: %v", cpf, err)
		return models.Customer{}, err
	}

	return customer, nil
}
