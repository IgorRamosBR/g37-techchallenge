package ports

import (
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/services/dto"
)

type CustomerService interface {
	GetCustomerById(id int) (domain.Customer, error)
	GetCustomerByCPF(cpf string) (domain.Customer, error)
	CreateCustomer(customerDTO dto.CustomerDTO) error
}

type ProductService interface {
	GetAllProducts(pageParameters dto.PageParams) (dto.Page[domain.Product], error)
	GetProductsByCategory(pageParameters dto.PageParams, category string) (dto.Page[domain.Product], error)
	GetProductById(id int) (domain.Product, error)
	CreateProduct(productDTO dto.ProductDTO) error
	UpdateProduct(id string, productDTO dto.ProductDTO) error
	DeleteProduct(id string) error
}

type OrderService interface {
	GetAllOrders(pageParameters dto.PageParams) (dto.Page[domain.Order], error)
	GetOrderStatus(orderId int) (dto.OrderStatusDTO, error)
	UpdateOrderStatus(orderId int, orderStatus string) error
	CreateOrder(orderDTO dto.OrderDTO) (dto.OrderCreationResponse, error)
}

type PaymentService interface {
	GeneratePaymentQRCode(order domain.Order) (string, error)
}
