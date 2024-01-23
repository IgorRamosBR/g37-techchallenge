package ports

import (
	"g37-lanchonete/internal/core/entities"
)

type PaymentOrderRepository interface {
	SavePaymentOrder(entities.PaymentOrder) error
}
