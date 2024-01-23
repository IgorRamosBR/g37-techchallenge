package ports

import (
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/services/dto"
)

type PaymentBroker interface {
	GeneratePaymentQRCode(dto.PaymentQRCodeRequest) (dto.PaymentQRCodeResponse, error)
}

type PaymentOrderRepository interface {
	SavePaymentOrder(domain.PaymentOrder) error
}
