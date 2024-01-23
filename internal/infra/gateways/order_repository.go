package gateways

import (
	"fmt"
	"g37-lanchonete/internal/core/entities"
	"g37-lanchonete/internal/core/usecases/dto"
	"g37-lanchonete/internal/infra/drivers/sql"
	"g37-lanchonete/internal/infra/gateways/sqlscripts"
)

type OrderRepositoryGateway interface {
	FindAllOrders(pageParams dto.PageParams) ([]entities.Order, error)
	GetOrderStatus(orderId int) (string, error)
	SaveOrder(order entities.Order) (int, error)
	UpdateOrderStatus(orderId int, orderStatus string) error
}

type orderRepositoryGateway struct {
	sqlClient sql.SQLClient
}

func NewOrderRepositoryGateway(sqlClient sql.SQLClient) OrderRepositoryGateway {
	return orderRepositoryGateway{
		sqlClient: sqlClient,
	}
}

func (r orderRepositoryGateway) FindAllOrders(pageParams dto.PageParams) ([]entities.Order, error) {
	rows, err := r.sqlClient.Find(sqlscripts.FindAllOrdersQuery, pageParams.GetLimit(), pageParams.GetOffset())
	if err != nil {
		return nil, fmt.Errorf("failed to find all orders, error %w", err)
	}

	orders := []entities.Order{}
	for rows.Next() {
		var order entities.Order
		var customer entities.Customer

		err = rows.Scan(&order.ID, &order.Coupon, &order.TotalAmount, &order.Status, &order.CreatedAt,
			&customer.ID, &customer.Name, &customer.Cpf, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan orders, error %w", err)
		}

		orderItems, err := r.getOrderItems(order.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order items, error %w", err)
		}

		order.Customer = customer
		order.Items = orderItems
		orders = append(orders, order)
	}

	return orders, nil
}

func (r orderRepositoryGateway) GetOrderStatus(orderId int) (string, error) {
	row := r.sqlClient.FindOne(sqlscripts.FindOrderStatusByIdQuery, orderId)

	var status string
	err := row.Scan(&status)
	if err != nil {
		return "", fmt.Errorf("failed to find order status, error %w", err)
	}

	return status, nil
}

func (r orderRepositoryGateway) SaveOrder(order entities.Order) (int, error) {
	tx, err := r.sqlClient.Begin()
	if err != nil {
		return -1, fmt.Errorf("failed to create a transaction, error %w", err)
	}

	row := tx.ExecWithReturn(sqlscripts.InsertOrderCmd, order.Coupon, order.TotalAmount, order.Customer.ID, order.Status, order.CreatedAt)

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

	err = tx.Commit()
	if err != nil {
		return -1, fmt.Errorf("failed to commit the transaction, error %w", err)
	}

	return orderId, nil
}

func (r orderRepositoryGateway) UpdateOrderStatus(orderId int, orderStatus string) error {
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

func (r orderRepositoryGateway) getOrderItems(orderId int) ([]entities.OrderItem, error) {
	rows, err := r.sqlClient.Find(sqlscripts.FindOrderItems, orderId)
	if err != nil {
		return nil, fmt.Errorf("failed to find order items, error %w", err)
	}

	orderItems := []entities.OrderItem{}
	for rows.Next() {
		var orderItem entities.OrderItem
		var product entities.Product

		err = rows.Scan(&orderItem.ID, &product.ID, &product.Name, &product.SkuId, &product.Description,
			&product.Category, &product.Price, &product.CreatedAt, &product.UpdatedAt, &orderItem.Quantity, &orderItem.Type)
		if err != nil {
			return nil, err
		}

		orderItem.Product = product
		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}
