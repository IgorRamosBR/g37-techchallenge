package ports

import (
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/services/dto"
)

type CustomerRepository interface {
	FindCustomerById(id int) (domain.Customer, error)
	FindCustomerByCPF(cpf string) (domain.Customer, error)
	SaveCustomer(customer domain.Customer) error
}

type ProductRepository interface {
	FindAllProducts(pageParams dto.PageParams) ([]domain.Product, error)
	FindProductsByCategory(pageParams dto.PageParams, category string) ([]domain.Product, error)
	FindProductById(id int) (domain.Product, error)
	SaveProduct(product domain.Product) error
	UpdateProduct(id int, product domain.Product) error
	DeleteProduct(id int) error
}

type OrderRepository interface {
	FindAllOrders(pageParams dto.PageParams) ([]domain.Order, error)
	SaveOrder(order domain.Order) error
	UpdateOrder(id uint, order domain.Order) error
}

type PaymentBroker interface {
	GeneratePaymentQRCode(dto.PaymentQRCodeRequest) (dto.PaymentQRCodeResponse, error)
}

type PaymentOrderRepository interface {
	SavePaymentOrder(domain.PaymentOrder) error
}
