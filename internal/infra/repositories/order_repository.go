package repositories

import (
	"errors"
	"fmt"
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/ports"
	"g37-lanchonete/internal/domain/services/dto"
	"g37-lanchonete/internal/infra/clients"
	"strconv"
)

type orderRepository struct {
	client clients.SQLClient
}

func NewOrderRepository(client clients.SQLClient) ports.OrderRepository {
	return orderRepository{
		client: client,
	}
}

func (r orderRepository) FindAllOrders(pageParams dto.PageParams) ([]models.Order, error) {
	var orders []models.Order
	err := r.client.FindAll(&orders, pageParams.GetLimit(), pageParams.GetOffset())
	if err != nil {
		return nil, fmt.Errorf("failed to find all orders, error %v", err)
	}

	return orders, nil
}

func (r orderRepository) SaveOrder(order models.Order) error {
	err := r.client.Save(&order)
	if err != nil {
		return fmt.Errorf("failed to save order, error %v", err)
	}

	return nil
}

func (r orderRepository) UpdateOrder(id uint, order models.Order) error {
	var oldOrder models.Order
	err := r.client.FindById(strconv.FormatUint(uint64(id), 10), &oldOrder)
	if err != nil {
		if errors.Is(err, clients.ErrNotFound) {
			return fmt.Errorf("order [%d] not found, error %v", id, err)
		}
		return fmt.Errorf("failed to find the order [%d], error %v", id, err)
	}

	order.ID = id
	err = r.client.Save(&order)
	if err != nil {
		return fmt.Errorf("failed to update the order, error %v", err)
	}

	return nil
}
