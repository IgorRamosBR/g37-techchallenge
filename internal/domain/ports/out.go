package ports

import "g37-lanchonete/internal/domain/models"

type CustomerRepository interface {
	SaveCustomer(customer models.Customer) error
	FindCustomerByCPF(cpf string) (models.Customer, error)
}
