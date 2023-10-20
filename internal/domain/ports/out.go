package ports

import "g37-lanchonete/internal/domain/models"

type CustomerRepository interface {
	SaveCustomer(customer models.Customer) error
	FindCustomerByCPF(cpf string) (models.Customer, error)
}

type ProductRepository interface {
	FindAllProducts() ([]models.Product, error)
	FindProductsByCategory(category string) ([]models.Product, error)
	SaveProduct(product models.Product) error
	UpdateProduct(id uint, product models.Product) error
	DeleteProduct(id uint) error
}
