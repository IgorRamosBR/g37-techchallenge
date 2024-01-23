package usecases

import (
	"g37-lanchonete/internal/core/entities"
	"g37-lanchonete/internal/core/usecases/dto"
	"g37-lanchonete/internal/infra/gateways"

	log "github.com/sirupsen/logrus"
)

type OrderUsecase interface {
	GetAllOrders(pageParameters dto.PageParams) (dto.Page[entities.Order], error)
	GetOrderStatus(orderId int) (dto.OrderStatusDTO, error)
	UpdateOrderStatus(orderId int, orderStatus string) error
	CreateOrder(orderDTO dto.OrderDTO) (dto.OrderCreationResponse, error)
}

type orderUsecase struct {
	customerUsecase        CustomerUsecase
	paymentUsecase         PaymentUsecase
	productUsecase         ProductUsecase
	orderRepositoryGateway gateways.OrderRepositoryGateway
}

func NewOrderUsecase(customerUsecase CustomerUsecase, paymentUsecase PaymentUsecase, productUsecase ProductUsecase, orderRepositoryGateway gateways.OrderRepositoryGateway) OrderUsecase {
	return orderUsecase{
		customerUsecase:        customerUsecase,
		paymentUsecase:         paymentUsecase,
		productUsecase:         productUsecase,
		orderRepositoryGateway: orderRepositoryGateway,
	}
}

func (u orderUsecase) GetAllOrders(pageParams dto.PageParams) (dto.Page[entities.Order], error) {
	orders, err := u.orderRepositoryGateway.FindAllOrders(pageParams)
	if err != nil {
		log.Errorf("failed to get all orders, error: %v", err)
		return dto.Page[entities.Order]{}, err
	}

	page := dto.BuildPage[entities.Order](orders, pageParams)
	return page, nil
}

func (u orderUsecase) CreateOrder(orderDTO dto.OrderDTO) (dto.OrderCreationResponse, error) {
	// Obter o cliente pelo ID
	customer, err := u.customerUsecase.GetCustomerById(orderDTO.CustomerId)
	if err != nil {
		log.Errorf("failed to get customer [%d], error: %v", orderDTO.CustomerId, err)
		return dto.OrderCreationResponse{}, err
	}

	// Criar um pedido a partir do DTO
	order := orderDTO.ToOrder(customer)

	// Calcular o total dos produtos
	totalAmount, err := u.calculateProducts(order.Items)
	if err != nil {
		log.Errorf("failed to calculate products, error: %v", err)
		return dto.OrderCreationResponse{}, err
	}

	// Definir o total no pedido
	order.TotalAmount = totalAmount

	// Salvar o pedido no banco de dados
	order.ID, err = u.saveOrder(order)
	if err != nil {
		log.Errorf("failed to save order, error: %v", err)
		return dto.OrderCreationResponse{}, err
	}

	// Gerar o código QR para o pagamento
	paymentQRCode, err := u.paymentUsecase.GeneratePaymentQRCode(order)
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

func (u orderUsecase) calculateProducts(items []entities.OrderItem) (float64, error) {
	for i, item := range items {
		product, err := u.getProduct(item.Product.ID)
		if err != nil {
			log.Errorf("failed to find products to process order, error: %v", err)
			return 0.0, err
		}
		item.Product = product
		items[i] = item
	}

	totalAmount := u.calculateTotal(items)
	return totalAmount, nil
}

func (u orderUsecase) getProduct(id int) (entities.Product, error) {
	product, err := u.productUsecase.GetProductById(id)
	if err != nil {
		log.Errorf("failed to find product [%d] to process order, error: %v", id, err)
		return entities.Product{}, err
	}

	return product, nil
}

func (u orderUsecase) calculateTotal(items []entities.OrderItem) float64 {
	var total float64
	for _, item := range items {
		total += item.Product.Price * float64(item.Quantity)
	}
	return total
}

func (u orderUsecase) saveOrder(order entities.Order) (int, error) {
	orderId, err := u.orderRepositoryGateway.SaveOrder(order)
	if err != nil {
		return 0, err
	}

	return orderId, nil
}

func (u orderUsecase) GetOrderStatus(orderId int) (dto.OrderStatusDTO, error) {
	status, err := u.orderRepositoryGateway.GetOrderStatus(orderId)
	if err != nil {
		return dto.OrderStatusDTO{}, err
	}

	return dto.OrderStatusDTO{
		Status: dto.OrderStatus(status),
	}, nil
}

func (u orderUsecase) UpdateOrderStatus(orderId int, orderStatus string) error {
	err := u.orderRepositoryGateway.UpdateOrderStatus(orderId, orderStatus)
	if err != nil {
		return err
	}

	return nil
}
