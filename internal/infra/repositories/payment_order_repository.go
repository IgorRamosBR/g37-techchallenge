package repositories

import (
	"encoding/json"
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/infra/clients"
)

type paymentOrderRepository struct {
	queue clients.Queue
}

func NewPaymentOrderRepository(queue clients.Queue) ports.PaymentOrderRepository {
	return paymentOrderRepository{
		queue: queue,
	}
}

func (r paymentOrderRepository) SavePaymentOrder(paymentOrder domain.PaymentOrder) error {
	payload, err := json.Marshal(paymentOrder)
	if err != nil {
		return err
	}

	err = r.queue.SendMessage(payload)
	if err != nil {
		return err
	}

	return nil
}
