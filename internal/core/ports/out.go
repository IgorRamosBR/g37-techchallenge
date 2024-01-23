package ports

import (
	"g37-lanchonete/internal/core/domain"
)

type PaymentOrderRepository interface {
	SavePaymentOrder(domain.PaymentOrder) error
}
