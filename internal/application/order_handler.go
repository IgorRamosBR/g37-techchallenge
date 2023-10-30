package application

import (
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService ports.OrderService
}

func NewOrderHandler(orderService ports.OrderService) OrderHandler {
	return OrderHandler{
		orderService: orderService,
	}
}

func (h OrderHandler) CreateOrder(c *gin.Context) {
	var order dto.OrderDTO
	err := c.ShouldBindJSON(&order)
	if err != nil {
		handleBadRequestResponse(c, "failed to bind order payload", err)
		return
	}

	valid, err := order.ValidateOrder()
	if !valid {
		handleBadRequestResponse(c, "invalid order payload", err)
		return
	}

	paymentQRCode, err := h.orderService.CreateOrder(order)
	if err != nil {
		handleInternalServerResponse(c, "failed to create product", err)
		return
	}

	c.JSON(http.StatusOK, dto.PaymentQRCode{QRCode: paymentQRCode})
}

func (h OrderHandler) GetAllOrders(c *gin.Context) {
	pageParams, err := getPageParams(c)
	if err != nil {
		handleBadRequestResponse(c, "invalid query parameters", err)
	}

	page, err := h.orderService.GetAllOrders(pageParams)
	if err != nil {
		handleInternalServerResponse(c, "failed to get all orders", err)
		return
	}

	c.JSON(http.StatusOK, page)
}
