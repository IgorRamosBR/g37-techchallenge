package dto

import (
	"g37-lanchonete/internal/domain/models"

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
	Product  ProductDTO
	Quantity int           `json:"quantity" valid:"int,required~Quantity is required|range(1|)~Quantity greater than 0"`
	Type     OrderItemType `json:"orderStatus" valid:"string,required~Type is required|in(UNIT|COMBO|CUSTOM_COMBO)~Type is invalid"`
}

func (o OrderItemDTO) toOrderItem() models.OrderItem {
	return models.OrderItem{
		Product:  o.Product.ToProduct(),
		Quantity: o.Quantity,
		Type:     string(o.Type),
	}
}

type OrderDTO struct {
	Items       []OrderItemDTO
	Coupon      string      `json:"coupon" valid:"length(0|100)~Description length should be less than 100 characters"`
	Discount    float64     `json:"discount" valid:"float,required~Discount is required|range(0.0|)~Discount must be posivitve"`
	CustomerId  string      `json:"customerId" valid:"length(0|200)~CustomerId length should be less than 200 characters"`
	OrderStatus OrderStatus `json:"orderStatus" valid:"string,required~OrderStatus is required|in(CREATED|PAID|RECEIVED|IN_PROGRESS|READY|DONE)~OrderStatus is invalid"`
}

func (o OrderDTO) ToOrder(customer models.Customer) models.Order {
	orderItems := make([]models.OrderItem, len(o.Items))
	for i, item := range o.Items {
		orderItems[i] = item.toOrderItem()
	}

	return models.Order{
		Items:    orderItems,
		Coupon:   o.Coupon,
		Discount: o.Discount,
		Customer: customer,
		Status:   string(o.OrderStatus),
	}
}

func (o OrderDTO) ValidateOrder() (bool, error) {
	if _, err := govalidator.ValidateStruct(o); err != nil {
		return false, err
	}

	return true, nil
}
