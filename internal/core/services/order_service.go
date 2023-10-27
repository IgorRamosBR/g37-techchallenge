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
	orderRepository ports.OrderRepository
}

func NewOrderService(customerService ports.CustomerService, paymentService ports.PaymentService, orderRepository ports.OrderRepository) ports.OrderService {
	return orderService{
		customerService: customerService,
		paymentService:  paymentService,
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
		log.Errorf("failed to get customer [%s], error: %v", orderDTO.CustomerId, err)
		return "", err
	}

	order := orderDTO.ToOrder(customer)
	err = s.saveOrder(order)
	if err != nil {
		log.Errorf("failed to save order, error: %v", err)
		return "", err
	}

	paymentQRCode, err := s.paymentService.ProcessPayment(order)
	if err != nil {
		log.Errorf("failed to process payment order, error: %v", err)
		return "", err
	}

	return paymentQRCode, nil
}

func (s orderService) saveOrder(order domain.Order) error {
	err := s.orderRepository.SaveOrder(order)
	if err != nil {
		return err
	}

	return nil
}
