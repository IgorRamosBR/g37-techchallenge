package ports

import (
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/services/dto"
)

type CustomerService interface {
	GetCustomerByCPF(cpf string) (models.Customer, error)
	CreateCustomer(customerDTO dto.CustomerDTO) error
}

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductsByCategory(category string) ([]models.Product, error)
	CreateProduct(productDTO dto.ProductDTO) error
	UpdateProduct(id string, productDTO dto.ProductDTO) error
	DeleteProduct(id string) error
}
