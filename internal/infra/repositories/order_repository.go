package repositories

import (
	"errors"
	"fmt"
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"g37-lanchonete/internal/infra/clients"
)

type orderRepository struct {
	client clients.SQLClient
}

func NewOrderRepository(client clients.SQLClient) ports.OrderRepository {
	return orderRepository{
		client: client,
	}
}

func (r orderRepository) FindAllOrders(pageParams dto.PageParams) ([]domain.Order, error) {
	var orders []domain.Order
	err := r.client.FindAll(&orders, pageParams.GetLimit(), pageParams.GetOffset(), []string{"Items", "Items.Products"})
	if err != nil {
		return nil, fmt.Errorf("failed to find all orders, error %v", err)
	}

	return orders, nil
}

func (r orderRepository) SaveOrder(order domain.Order) error {
	err := r.client.Save(&order)
	if err != nil {
		return fmt.Errorf("failed to save order, error %v", err)
	}

	for _, item := range order.Items {
		err := r.client.SaveAssociations(&item, "Products", item.Products)
		if err != nil {
			return fmt.Errorf("failed to save order items associations, error %v", err)
		}
	}

	return nil
}

func (r orderRepository) UpdateOrder(id uint, order domain.Order) error {
	var oldOrder domain.Order
	err := r.client.FindById(int(id), &oldOrder)
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
