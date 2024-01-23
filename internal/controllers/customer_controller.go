package controllers

import (
	"errors"
	"g37-lanchonete/internal/core/usecases"
	"g37-lanchonete/internal/core/usecases/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomeController struct {
	customerUsecase usecases.CustomerUsecase
}

func NewCustomerController(customerUsecase usecases.CustomerUsecase) CustomeController {
	return CustomeController{
		customerUsecase: customerUsecase,
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

	err = c.customerUsecase.CreateCustomer(customer)
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

	customer, err := c.customerUsecase.GetCustomerByCPF(cpf)
	if err != nil {
		handleNotFoundRequestResponse(ctx, "failed to find customer", err)
		return
	}

	ctx.JSON(200, customer)
}
