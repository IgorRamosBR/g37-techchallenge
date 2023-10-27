package services

import (
	"fmt"
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type PaymentService interface {
	ProcessPayment(order domain.Order) (string, error)
}

type paymentService struct {
	notificationUrl        string
	sponsorId              string
	paymentBroker          ports.PaymentBroker
	paymentOrderRepository ports.PaymentOrderRepository
}

func NewPaymentService(notificationUrl, sponsorId string, paymentBroker ports.PaymentBroker, paymentOrderRepository ports.PaymentOrderRepository) PaymentService {
	return paymentService{
		notificationUrl:        notificationUrl,
		sponsorId:              sponsorId,
		paymentBroker:          paymentBroker,
		paymentOrderRepository: paymentOrderRepository,
	}
}

func (p paymentService) ProcessPayment(order domain.Order) (string, error) {
	qrcode, err := p.generatePaymentQRCode(order)
	if err != nil {
		log.Errorf("failed to generate payment qrcode for the order [%d], error: %v", order.ID, err)
		return "", err
	}

	paymentOrder := domain.PaymentOrder{Order: strconv.FormatUint(uint64(order.ID), 10)}
	err = p.paymentOrderRepository.SavePaymentOrder(paymentOrder)
	if err != nil {
		log.Errorf("failed to save payment order for the order [%d], error: %v", order.ID, err)
		return "", err
	}

	return qrcode, nil
}

func (p paymentService) generatePaymentQRCode(order domain.Order) (string, error) {
	paymentRequest := p.createPaymentRequest(order)
	paymentResponse, err := p.paymentBroker.GeneratePaymentQRCode(paymentRequest)
	if err != nil {
		return "", err
	}

	return paymentResponse.QrData, nil
}

func (p paymentService) createPaymentRequest(order domain.Order) dto.PaymentQRCodeRequest {
	items := make([]dto.PaymentItemRequest, len(order.Items))
	for i, item := range order.Items {
		items[i] = createPaymentItem(item)
	}

	return dto.PaymentQRCodeRequest{
		ExternalReference: strconv.FormatUint(uint64(order.ID), 10),
		Title:             fmt.Sprintf("Order %d for the Customer[%d]", order.ID, order.CustomerID),
		NotificationURL:   p.notificationUrl,
		TotalAmount:       order.TotalAmount,
		Items:             items,
		Sponsor:           p.sponsorId,
	}
}

func createPaymentItem(item domain.OrderItem) dto.PaymentItemRequest {
	return dto.PaymentItemRequest{
		SkuNumber:   item.Product.SkuId,
		Category:    item.Product.Category,
		Title:       item.Product.Name,
		Description: item.Product.Description,
		UnitPrice:   item.Product.Price,
		Quantity:    item.Quantity,
		UnitMeasure: "unit",
		TotalAmount: item.Product.Price * float64(item.Quantity),
	}
}
