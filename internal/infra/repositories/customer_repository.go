package repositories

import (
	"fmt"
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/infra/clients/sql"
	"g37-lanchonete/internal/infra/sqlscripts"
)

type customerRepository struct {
	sqlClient sql.SQLClient
}

func NewCustomerRepository(sqlClient sql.SQLClient) ports.CustomerRepository {
	return customerRepository{
		sqlClient: sqlClient,
	}
}

func (r customerRepository) FindCustomerById(id int) (domain.Customer, error) {
	getCustomerByIdQuery := fmt.Sprintf(sqlscripts.GetCustomerByIdQuery)

	row := r.sqlClient.FindOne(getCustomerByIdQuery, id)

	var customer domain.Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.Cpf, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return domain.Customer{}, fmt.Errorf("failed to find customer by id [%d], error %v", id, err)
	}

	return customer, nil
}

func (r customerRepository) FindCustomerByCPF(cpf string) (domain.Customer, error) {
	getCustomerByIdQuery := fmt.Sprintf(sqlscripts.GetCustomerByIdQuery)

	row := r.sqlClient.FindOne(getCustomerByIdQuery, cpf)

	var customer domain.Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.Cpf, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return domain.Customer{}, fmt.Errorf("failed to find customer by cpf [%s], error %v", cpf, err)
	}

	return customer, nil
}

func (r customerRepository) SaveCustomer(customer domain.Customer) error {
	insertCustomerCmd := fmt.Sprintf(sqlscripts.InsertCustomer)

	_, err := r.sqlClient.Exec(insertCustomerCmd, customer.Name, customer.Cpf, customer.Email, customer.CreatedAt, customer.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save customer, error %v", err)
	}

	return nil
}
