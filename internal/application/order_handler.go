package application

import (
	"errors"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"net/http"
	"strconv"

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

	createResponse, err := h.orderService.CreateOrder(order)
	if err != nil {
		handleInternalServerResponse(c, "failed to create product", err)
		return
	}

	c.JSON(http.StatusOK, dto.OrderCreationResponse{QRCode: createResponse.QRCode, OrderID: createResponse.OrderID})
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

func (h OrderHandler) GetOrderStatus(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handleBadRequestResponse(c, "[id] path parameter is required", errors.New("id is missing"))
		return
	}

	orderID, err := strconv.Atoi(id)
	if err != nil {
		handleBadRequestResponse(c, "[id] path parameter is invalid", err)
		return
	}

	response, err := h.orderService.GetOrderStatus(orderID)
	if err != nil {
		handleInternalServerResponse(c, "failed to get order status", err)
		return
	}

	c.JSON(http.StatusOK, response)

}

func (h OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handleBadRequestResponse(c, "[id] path parameter is required", errors.New("id is missing"))
		return
	}

	orderId, err := strconv.Atoi(id)
	if err != nil {
		handleBadRequestResponse(c, "[id] path parameter is invalid", err)
		return
	}

	var orderStatus dto.OrderStatusDTO
	err = c.ShouldBindJSON(&orderStatus)
	if err != nil {
		handleBadRequestResponse(c, "failed to bind order status payload", err)
		return
	}

	valid, err := orderStatus.Validate()
	if !valid {
		handleBadRequestResponse(c, "invalid order status payload", err)
		return
	}

	err = h.orderService.UpdateOrderStatus(orderId, string(orderStatus.Status))
	if err != nil {
		handleInternalServerResponse(c, "failed to update order status", err)
		return
	}

	c.Status(http.StatusNoContent)
}
