package gateways

import (
	"fmt"
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/infra/clients/sql"
	"g37-lanchonete/internal/infra/sqlscripts"
)

type CustomerRepositoryGateway interface {
	FindCustomerById(id int) (domain.Customer, error)
	FindCustomerByCPF(cpf string) (domain.Customer, error)
	SaveCustomer(customer domain.Customer) error
}

type customerRepositoryGateway struct {
	sqlClient sql.SQLClient
}

func NewCustomerRepositoryGateway(sqlClient sql.SQLClient) CustomerRepositoryGateway {
	return customerRepositoryGateway{
		sqlClient: sqlClient,
	}
}

func (r customerRepositoryGateway) FindCustomerById(id int) (domain.Customer, error) {
	getCustomerByIdQuery := fmt.Sprintf(sqlscripts.GetCustomerByIdQuery)

	row := r.sqlClient.FindOne(getCustomerByIdQuery, id)

	var customer domain.Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.Cpf, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return domain.Customer{}, fmt.Errorf("failed to find customer by id [%d], error %v", id, err)
	}

	return customer, nil
}

func (r customerRepositoryGateway) FindCustomerByCPF(cpf string) (domain.Customer, error) {
	getCustomerByIdQuery := fmt.Sprintf(sqlscripts.GetCustomerByCPFQuery)

	row := r.sqlClient.FindOne(getCustomerByIdQuery, cpf)

	var customer domain.Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.Cpf, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return domain.Customer{}, fmt.Errorf("failed to find customer by cpf [%s], error %v", cpf, err)
	}

	return customer, nil
}

func (r customerRepositoryGateway) SaveCustomer(customer domain.Customer) error {
	insertCustomerCmd := fmt.Sprintf(sqlscripts.InsertCustomer)

	_, err := r.sqlClient.Exec(insertCustomerCmd, customer.Name, customer.Cpf, customer.Email, customer.CreatedAt, customer.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save customer, error %v", err)
	}

	return nil
}
