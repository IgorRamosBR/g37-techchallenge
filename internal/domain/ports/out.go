package ports

import (
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/services/dto"
)

type CustomerRepository interface {
	FindCustomerById(id string) (models.Customer, error)
	FindCustomerByCPF(cpf string) (models.Customer, error)
	SaveCustomer(customer models.Customer) error
}

type ProductRepository interface {
	FindAllProducts(pageParams dto.PageParams) ([]models.Product, error)
	FindProductsByCategory(pageParams dto.PageParams, category string) ([]models.Product, error)
	SaveProduct(product models.Product) error
	UpdateProduct(id uint, product models.Product) error
	DeleteProduct(id uint) error
}

type OrderRepository interface {
	FindAllOrders(pageParams dto.PageParams) ([]models.Order, error)
	SaveOrder(order models.Order) error
	UpdateOrder(id uint, order models.Order) error
}

type PaymentBroker interface {
	GeneratePaymentQRCode(dto.PaymentQRCodeRequest) (dto.PaymentQRCodeResponse, error)
}

type PaymentOrderRepository interface {
	SavePaymentOrder(models.PaymentOrder) error
}
