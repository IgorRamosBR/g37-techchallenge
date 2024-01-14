package dto

import (
	"g37-lanchonete/internal/core/domain"
	"time"

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

type OrderStatusRequest struct {
	Status OrderStatus `json:"status" valid:"in(CREATED|PAID|RECEIVED,IN_PROGRESS,READY,DONE),required~Status is invalid"`
}

func (o OrderStatusRequest) Validate() (bool, error) {
	if _, err := govalidator.ValidateStruct(o); err != nil {
		return false, err
	}

	return true, nil
}

type OrderItemType string

const (
	OrderItemTypeUnit        OrderItemType = "UNIT"
	OrderItemTypeCombo       OrderItemType = "COMBO"
	OrderItemTypeCustomCombo OrderItemType = "CUSTOM_COMBO"
)

type OrderItemDTO struct {
	ProductId int           `json:"productIds"`
	Quantity  int           `json:"quantity" valid:"int,required~Quantity is required|range(1|)~Quantity greater than 0"`
	Type      OrderItemType `json:"type" valid:"in(UNIT|COMBO|CUSTOM_COMBO),required~Type is invalid"`
}

func (o OrderItemDTO) toOrderItem() domain.OrderItem {
	return domain.OrderItem{
		Product: domain.Product{
			ID: o.ProductId,
		},
		Quantity: o.Quantity,
		Type:     string(o.Type),
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
		Items:     orderItems,
		Coupon:    o.Coupon,
		Customer:  customer,
		Status:    string(o.Status),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (o OrderDTO) ValidateOrder() (bool, error) {
	if _, err := govalidator.ValidateStruct(o); err != nil {
		return false, err
	}

	return true, nil
}
