package usecases

import (
	"fmt"
	"g37-lanchonete/internal/core/entities"
	"g37-lanchonete/internal/core/usecases/dto"
	"g37-lanchonete/internal/infra/drivers/payment"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type PaymentUsecase interface {
	GeneratePaymentQRCode(order entities.Order) (string, error)
}

type paymentUsecase struct {
	notificationUrl string
	sponsorId       string
	paymentBroker   payment.PaymentBroker
}

func NewPaymentUsecase(notificationUrl, sponsorId string, paymentBroker payment.PaymentBroker) PaymentUsecase {
	return paymentUsecase{
		notificationUrl: notificationUrl,
		sponsorId:       sponsorId,
		paymentBroker:   paymentBroker,
	}
}

func (u paymentUsecase) GeneratePaymentQRCode(order entities.Order) (string, error) {
	paymentRequest := u.createPaymentRequest(order)
	paymentResponse, err := u.paymentBroker.GeneratePaymentQRCode(paymentRequest)
	if err != nil {
		log.Errorf("failed to generate payment qrcode for the order [%d], error: %v", order.ID, err)
		return "", err
	}

	return paymentResponse.QrData, nil
}

func (u paymentUsecase) createPaymentRequest(order entities.Order) dto.PaymentQRCodeRequest {
	var items []dto.PaymentItemRequest
	for _, item := range order.Items {
		items = append(items, createPaymentItem(item))
	}

	return dto.PaymentQRCodeRequest{
		ExternalReference: strconv.FormatUint(uint64(order.ID), 10),
		Title:             fmt.Sprintf("Order %d for the Customer[%d]", order.ID, order.Customer.ID),
		NotificationURL:   fmt.Sprintf("%s/orders/%d/payment", u.notificationUrl, order.ID),
		TotalAmount:       order.TotalAmount,
		Items:             items,
		Sponsor:           u.sponsorId,
	}
}

func createPaymentItem(item entities.OrderItem) dto.PaymentItemRequest {
	paymentItem := dto.PaymentItemRequest{
		SkuNumber:   item.Product.SkuId,
		Category:    item.Product.Category,
		Title:       item.Product.Name,
		Description: item.Product.Description,
		UnitPrice:   item.Product.Price,
		Quantity:    item.Quantity,
		UnitMeasure: getUnitMeasure(item.Type),
		TotalAmount: item.Product.Price * float64(item.Quantity),
	}

	return paymentItem
}

func getUnitMeasure(itemType string) string {
	if itemType == string(dto.OrderItemTypeCustomCombo) {
		return "pack"
	}
	return "unit"
}
