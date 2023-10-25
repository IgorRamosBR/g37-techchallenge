package ports

import (
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/services/dto"
)

type CustomerService interface {
	GetCustomerById(id string) (models.Customer, error)
	GetCustomerByCPF(cpf string) (models.Customer, error)
	CreateCustomer(customerDTO dto.CustomerDTO) error
}

type ProductService interface {
	GetAllProducts(pageParameters dto.PageParams) (dto.Page[models.Product], error)
	GetProductsByCategory(pageParameters dto.PageParams, category string) (dto.Page[models.Product], error)
	CreateProduct(productDTO dto.ProductDTO) error
	UpdateProduct(id string, productDTO dto.ProductDTO) error
	DeleteProduct(id string) error
}

type OrderService interface {
	GetAllOrders(pageParameters dto.PageParams) (dto.Page[models.Order], error)
	CreateOrder(orderDTO dto.OrderDTO) (string, error)
}

type PaymentService interface {
	ProcessPayment(order models.Order) (string, error)
}
