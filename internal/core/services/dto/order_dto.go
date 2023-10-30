package dto

import (
	"fmt"
	"g37-lanchonete/internal/core/domain"

	"github.com/asaskevich/govalidator"
)

type OrderStatus string

const (
	OrderStatusCreated    OrderStatus = "CREATED"
	OrderStatusPaid       OrderStatus = "PAID"
	OrderStatusReceived   OrderStatus = "RECEIVED"
	OrderStatusInProgress OrderStatus = "IN_PROGRESS"
	OrderStatusReady      OrderStatus = "READY"
	OrderStatusDone       OrderStatus = "DONE"
)

type OrderItemType string

const (
	OrderItemTypeUnit        OrderItemType = "UNIT"
	OrderItemTypeCombo       OrderItemType = "COMBO"
	OrderItemTypeCustomCombo OrderItemType = "CUSTOM_COMBO"
)

type OrderItemDTO struct {
	ProductIds []int         `json:"productIds"`
	Quantity   int           `json:"quantity" valid:"int,required~Quantity is required|range(1|)~Quantity greater than 0"`
	Type       OrderItemType `json:"type" valid:"in(UNIT|COMBO|CUSTOM_COMBO),required~Type is invalid"`
}

func (o OrderItemDTO) toOrderItem() domain.OrderItem {
	return domain.OrderItem{
		ProductIds: o.ProductIds,
		Quantity:   o.Quantity,
		Type:       string(o.Type),
	}
}

type OrderDTO struct {
	Items      []OrderItemDTO `json:"items"`
	Coupon     string         `json:"coupon" valid:"length(0|100)~Description length should be less than 100 characters"`
	CustomerId int            `json:"customerId" valid:"length(0|200)~CustomerId length should be less than 200 characters"`
	Status     OrderStatus    `json:"status" valid:"in(CREATED|PAID|RECEIVED|IN_PROGRESS|READY|DONE),required~Status is invalid"`
}

func (o OrderDTO) ToOrder(customer domain.Customer) domain.Order {
	orderItems := make([]domain.OrderItem, len(o.Items))
	for i, item := range o.Items {
		orderItems[i] = item.toOrderItem()
	}

	return domain.Order{
		Items:      orderItems,
		Coupon:     o.Coupon,
		CustomerID: customer.ID,
		Status:     string(o.Status),
	}
}

func (o OrderDTO) ValidateOrder() (bool, error) {
	if _, err := govalidator.ValidateStruct(o); err != nil {
		return false, err
	}

	for _, item := range o.Items {
		if item.Type != OrderItemTypeCustomCombo && len(item.ProductIds) > 1 {
			return false, fmt.Errorf("item type [%s] must have only one productId", item.Type)
		}
	}

	return true, nil
}
