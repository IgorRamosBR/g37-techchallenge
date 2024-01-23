package services

import (
	"g37-lanchonete/internal/core/entities"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"g37-lanchonete/internal/infra/gateways"

	log "github.com/sirupsen/logrus"
)

type orderService struct {
	customerService        ports.CustomerService
	paymentService         ports.PaymentService
	productService         ports.ProductService
	orderRepositoryGateway gateways.OrderRepositoryGateway
}

func NewOrderService(customerService ports.CustomerService, paymentService ports.PaymentService, productService ports.ProductService, orderRepositoryGateway gateways.OrderRepositoryGateway) ports.OrderService {
	return orderService{
		customerService:        customerService,
		paymentService:         paymentService,
		productService:         productService,
		orderRepositoryGateway: orderRepositoryGateway,
	}
}

func (s orderService) GetAllOrders(pageParams dto.PageParams) (dto.Page[entities.Order], error) {
	orders, err := s.orderRepositoryGateway.FindAllOrders(pageParams)
	if err != nil {
		log.Errorf("failed to get all orders, error: %v", err)
		return dto.Page[entities.Order]{}, err
	}

	page := dto.BuildPage[entities.Order](orders, pageParams)
	return page, nil
}

func (s orderService) CreateOrder(orderDTO dto.OrderDTO) (dto.OrderCreationResponse, error) {
	// Obter o cliente pelo ID
	customer, err := s.customerService.GetCustomerById(orderDTO.CustomerId)
	if err != nil {
		log.Errorf("failed to get customer [%d], error: %v", orderDTO.CustomerId, err)
		return dto.OrderCreationResponse{}, err
	}

	// Criar um pedido a partir do DTO
	order := orderDTO.ToOrder(customer)

	// Calcular o total dos produtos
	totalAmount, err := s.calculateProducts(order.Items)
	if err != nil {
		log.Errorf("failed to calculate products, error: %v", err)
		return dto.OrderCreationResponse{}, err
	}

	// Definir o total no pedido
	order.TotalAmount = totalAmount

	// Salvar o pedido no banco de dados
	order.ID, err = s.saveOrder(order)
	if err != nil {
		log.Errorf("failed to save order, error: %v", err)
		return dto.OrderCreationResponse{}, err
	}

	// Gerar o código QR para o pagamento
	paymentQRCode, err := s.paymentService.GeneratePaymentQRCode(order)
	if err != nil {
		log.Errorf("failed to process payment order, error: %v", err)
		return dto.OrderCreationResponse{}, err
	}

	// Construir a resposta com o código QR e o ID do pedido
	response := dto.OrderCreationResponse{
		QRCode:  paymentQRCode,
		OrderID: order.ID,
	}

	return response, nil
}

func (s orderService) calculateProducts(items []entities.OrderItem) (float64, error) {
	for i, item := range items {
		product, err := s.getProduct(item.Product.ID)
		if err != nil {
			log.Errorf("failed to find products to process order, error: %v", err)
			return 0.0, err
		}
		item.Product = product
		items[i] = item
	}

	totalAmount := s.calculateTotal(items)
	return totalAmount, nil
}

func (s orderService) getProduct(id int) (entities.Product, error) {
	product, err := s.productService.GetProductById(id)
	if err != nil {
		log.Errorf("failed to find product [%d] to process order, error: %v", id, err)
		return entities.Product{}, err
	}

	return product, nil
}

func (s orderService) calculateTotal(items []entities.OrderItem) float64 {
	var total float64
	for _, item := range items {
		total += item.Product.Price * float64(item.Quantity)
	}
	return total
}

func (s orderService) saveOrder(order entities.Order) (int, error) {
	orderId, err := s.orderRepositoryGateway.SaveOrder(order)
	if err != nil {
		return 0, err
	}

	return orderId, nil
}

func (s orderService) GetOrderStatus(orderId int) (dto.OrderStatusDTO, error) {
	status, err := s.orderRepositoryGateway.GetOrderStatus(orderId)
	if err != nil {
		return dto.OrderStatusDTO{}, err
	}

	return dto.OrderStatusDTO{
		Status: dto.OrderStatus(status),
	}, nil
}

func (s orderService) UpdateOrderStatus(orderId int, orderStatus string) error {
	err := s.orderRepositoryGateway.UpdateOrderStatus(orderId, orderStatus)
	if err != nil {
		return err
	}

	return nil
}
