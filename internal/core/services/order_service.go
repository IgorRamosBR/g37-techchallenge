package services

import (
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"

	log "github.com/sirupsen/logrus"
)

type orderService struct {
	customerService ports.CustomerService
	paymentService  ports.PaymentService
	productService  ports.ProductService
	orderRepository ports.OrderRepository
}

func NewOrderService(customerService ports.CustomerService, paymentService ports.PaymentService, productService ports.ProductService, orderRepository ports.OrderRepository) ports.OrderService {
	return orderService{
		customerService: customerService,
		paymentService:  paymentService,
		productService:  productService,
		orderRepository: orderRepository,
	}
}

func (s orderService) GetAllOrders(pageParams dto.PageParams) (dto.Page[domain.Order], error) {
	orders, err := s.orderRepository.FindAllOrders(pageParams)
	if err != nil {
		log.Errorf("failed to get all orders, error: %v", err)
		return dto.Page[domain.Order]{}, err
	}

	page := dto.BuildPage[domain.Order](orders, pageParams)
	return page, nil
}

func (s orderService) CreateOrder(orderDTO dto.OrderDTO) (string, error) {
	customer, err := s.customerService.GetCustomerById(orderDTO.CustomerId)
	if err != nil {
		log.Errorf("failed to get customer [%d], error: %v", orderDTO.CustomerId, err)
		return "", err
	}

	order := orderDTO.ToOrder(customer)

	totalAmout, err := s.calculateProducts(order.Items)
	if err != nil {
		log.Errorf("failed to calculate products, error: %v", err)
		return "", err
	}

	order.TotalAmount = totalAmout
	err = s.saveOrder(order)
	if err != nil {
		log.Errorf("failed to save order, error: %v", err)
		return "", err
	}

	paymentQRCode, err := s.paymentService.GeneratePaymentQRCode(order)
	if err != nil {
		log.Errorf("failed to process payment order, error: %v", err)
		return "", err
	}

	return paymentQRCode, nil
}

func (s orderService) calculateProducts(items []domain.OrderItem) (float64, error) {
	for i, item := range items {
		product, err := s.productService.GetProductById(item.ProductID)
		if err != nil {
			log.Errorf("failed to find product to process order, error: %v", err)
			return 0.0, err
		}
		item.Product = product
		items[i] = item
	}

	totalAmount := s.calculateTotal(items)
	return totalAmount, nil
}

func (s orderService) calculateTotal(items []domain.OrderItem) float64 {
	var total float64
	for _, item := range items {
		total += item.Product.Price * float64(item.Quantity)
	}

	return total
}

func (s orderService) saveOrder(order domain.Order) error {
	err := s.orderRepository.SaveOrder(order)
	if err != nil {
		return err
	}

	return nil
}
