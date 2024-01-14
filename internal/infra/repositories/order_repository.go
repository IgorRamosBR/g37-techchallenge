package repositories

import (
	"fmt"
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"g37-lanchonete/internal/infra/clients/sql"
	"g37-lanchonete/internal/infra/sqlscripts"
)

type orderRepository struct {
	sqlClient sql.SQLClient
}

func NewOrderRepository(sqlClient sql.SQLClient) ports.OrderRepository {
	return orderRepository{
		sqlClient: sqlClient,
	}
}

func (r orderRepository) FindAllOrders(pageParams dto.PageParams) ([]domain.Order, error) {
	rows, err := r.sqlClient.Find(sqlscripts.FindAllOrdersQuery, pageParams.GetLimit(), pageParams.GetOffset())
	if err != nil {
		return nil, fmt.Errorf("failed to find all orders, error %w", err)
	}

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		var customer domain.Customer

		err = rows.Scan(&order.ID, &order.Coupon, &order.TotalAmount, &order.Status, &order.CreatedAt,
			&customer.ID, &customer.Name, &customer.Cpf, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan orders, error %w", err)
		}

		order.Customer = customer
		orders = append(orders, order)
	}

	return orders, nil
}

func (r orderRepository) GetOrderStatus(orderId int) (string, error) {
	row := r.sqlClient.FindOne(sqlscripts.FindOrderStatusByIdQuery, orderId)

	var status string
	err := row.Scan(&status)
	if err != nil {
		return "", fmt.Errorf("failed to find order status, error %w", err)
	}

	return status, nil
}

func (r orderRepository) SaveOrder(order domain.Order) (int, error) {
	tx, err := r.sqlClient.Begin()
	if err != nil {
		return -1, fmt.Errorf("failed to create a transaction, error %w", err)
	}

	row := tx.ExecWithReturn(sqlscripts.InsertOrderCmd, order.Coupon, order.TotalAmount, order.Customer.ID, order.Status, order.CreatedAt, order.UpdatedAt)

	var orderId int
	err = row.Scan(&orderId)
	if err != nil {
		return -1, fmt.Errorf("failed to save order, error %w", err)
	}

	for _, item := range order.Items {
		_, err := tx.Exec(sqlscripts.InsertOrderItemCmd, orderId, item.Product.ID, item.Quantity, item.Type)
		if err != nil {
			return -1, fmt.Errorf("failed to save order items associations, error %v", err)
		}
	}

	return orderId, nil
}

func (r orderRepository) UpdateOrderStatus(orderId int, orderStatus string) error {
	result, err := r.sqlClient.Exec(sqlscripts.UpdateOrderStatusCmd, orderId, orderStatus)
	if err != nil {
		return fmt.Errorf("failed to update order status, error %w", err)
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check order status update operation, error %w", err)
	}

	if rowsAffect < 1 {
		return sql.ErrNotFound
	}

	return nil
}
