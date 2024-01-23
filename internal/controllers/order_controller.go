package controllers

import (
	"errors"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService ports.OrderService
}

func NewOrderController(orderService ports.OrderService) OrderController {
	return OrderController{
		orderService: orderService,
	}
}

func (c OrderController) CreateOrder(ctx *gin.Context) {
	var order dto.OrderDTO
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		handleBadRequestResponse(ctx, "failed to bind order payload", err)
		return
	}

	valid, err := order.ValidateOrder()
	if !valid {
		handleBadRequestResponse(ctx, "invalid order payload", err)
		return
	}

	createResponse, err := c.orderService.CreateOrder(order)
	if err != nil {
		handleInternalServerResponse(ctx, "failed to create product", err)
		return
	}

	ctx.JSON(http.StatusOK, dto.OrderCreationResponse{QRCode: createResponse.QRCode, OrderID: createResponse.OrderID})
}

func (c OrderController) GetAllOrders(ctx *gin.Context) {
	pageParams, err := getPageParams(ctx)
	if err != nil {
		handleBadRequestResponse(ctx, "invalid query parameters", err)
	}

	page, err := c.orderService.GetAllOrders(pageParams)
	if err != nil {
		handleInternalServerResponse(ctx, "failed to get all orders", err)
		return
	}

	ctx.JSON(http.StatusOK, page)
}

func (c OrderController) GetOrderStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		handleBadRequestResponse(ctx, "[id] path parameter is required", errors.New("id is missing"))
		return
	}

	orderID, err := strconv.Atoi(id)
	if err != nil {
		handleBadRequestResponse(ctx, "[id] path parameter is invalid", err)
		return
	}

	response, err := c.orderService.GetOrderStatus(orderID)
	if err != nil {
		handleInternalServerResponse(ctx, "failed to get order status", err)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (c OrderController) UpdateOrderStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		handleBadRequestResponse(ctx, "[id] path parameter is required", errors.New("id is missing"))
		return
	}

	orderId, err := strconv.Atoi(id)
	if err != nil {
		handleBadRequestResponse(ctx, "[id] path parameter is invalid", err)
		return
	}

	var orderStatus dto.OrderStatusDTO
	err = ctx.ShouldBindJSON(&orderStatus)
	if err != nil {
		handleBadRequestResponse(ctx, "failed to bind order status payload", err)
		return
	}

	valid, err := orderStatus.Validate()
	if !valid {
		handleBadRequestResponse(ctx, "invalid order status payload", err)
		return
	}

	err = c.orderService.UpdateOrderStatus(orderId, string(orderStatus.Status))
	if err != nil {
		handleInternalServerResponse(ctx, "failed to update order status", err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
