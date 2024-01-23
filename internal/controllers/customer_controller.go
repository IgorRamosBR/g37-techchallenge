package controllers

import (
	"errors"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomeController struct {
	customerService ports.CustomerService
}

func NewCustomerController(customerService ports.CustomerService) CustomeController {
	return CustomeController{
		customerService: customerService,
	}
}

func (c CustomeController) SaveCustomer(ctx *gin.Context) {
	var customer dto.CustomerDTO
	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		handleBadRequestResponse(ctx, "failed to bind customer payload", err)
		return
	}

	valid, err := customer.ValidateCustomer()
	if !valid {
		handleBadRequestResponse(ctx, "invalid customer payload", err)
		return
	}

	err = c.customerService.CreateCustomer(customer)
	if err != nil {
		handleInternalServerResponse(ctx, "failed to create customer", err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c CustomeController) GetCustomers(ctx *gin.Context) {
	cpf := ctx.Query("cpf")
	if cpf == "" {
		handleBadRequestResponse(ctx, "cpf query parameter is required", errors.New("cpf is missing"))
		return
	}

	customer, err := c.customerService.GetCustomerByCPF(cpf)
	if err != nil {
		handleNotFoundRequestResponse(ctx, "failed to find customer", err)
		return
	}

	ctx.JSON(200, customer)
}
