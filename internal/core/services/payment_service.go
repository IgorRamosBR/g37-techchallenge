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
	GeneratePaymentQRCode(order domain.Order) (string, error)
}

type paymentService struct {
	notificationUrl string
	sponsorId       string
	paymentBroker   ports.PaymentBroker
}

func NewPaymentService(notificationUrl, sponsorId string, paymentBroker ports.PaymentBroker) PaymentService {
	return paymentService{
		notificationUrl: notificationUrl,
		sponsorId:       sponsorId,
		paymentBroker:   paymentBroker,
	}
}

func (p paymentService) GeneratePaymentQRCode(order domain.Order) (string, error) {
	paymentRequest := p.createPaymentRequest(order)
	paymentResponse, err := p.paymentBroker.GeneratePaymentQRCode(paymentRequest)
	if err != nil {
		log.Errorf("failed to generate payment qrcode for the order [%d], error: %v", order.ID, err)
		return "", err
	}

	return paymentResponse.QrData, nil
}

func (p paymentService) createPaymentRequest(order domain.Order) dto.PaymentQRCodeRequest {
	var items []dto.PaymentItemRequest
	for _, item := range order.Items {
		items = append(items, createPaymentItem(item)...)
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

func createPaymentItem(item domain.OrderItem) []dto.PaymentItemRequest {
	paymentItems := make([]dto.PaymentItemRequest, len(item.Products))
	for i, product := range item.Products {
		paymentItem := dto.PaymentItemRequest{
			SkuNumber:   product.SkuId,
			Category:    product.Category,
			Title:       product.Name,
			Description: product.Description,
			UnitPrice:   product.Price,
			Quantity:    item.Quantity,
			UnitMeasure: getUnitMeasure(item.Type),
			TotalAmount: product.Price * float64(item.Quantity),
		}
		paymentItems[i] = paymentItem
	}

	return paymentItems
}

func getUnitMeasure(itemType string) string {
	if itemType == string(dto.OrderItemTypeCustomCombo) {
		return "pack"
	}
	return "unit"
}
