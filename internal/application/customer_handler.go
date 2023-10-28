package application

import (
	"errors"
	"g37-lanchonete/internal/core/ports"
	"g37-lanchonete/internal/core/services/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService ports.CustomerService
}

func NewCustomerHandler(customerService ports.CustomerService) CustomerHandler {
	return CustomerHandler{
		customerService: customerService,
	}
}

func (h CustomerHandler) SaveCustomer(c *gin.Context) {
	var customer dto.CustomerDTO
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		handleBadRequestResponse(c, "failed to bind customer payload", err)
		return
	}

	valid, err := customer.ValidateCustomer()
	if !valid {
		handleBadRequestResponse(c, "invalid customer payload", err)
		return
	}

	err = h.customerService.CreateCustomer(customer)
	if err != nil {
		handleInternalServerResponse(c, "failed to create customer", err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h CustomerHandler) GetCustomers(c *gin.Context) {
	cpf := c.Query("cpf")
	if cpf == "" {
		handleBadRequestResponse(c, "cpf query parameter is required", errors.New("cpf is missing"))
		return
	}

	customer, err := h.customerService.GetCustomerByCPF(cpf)
	if err != nil {
		handleNotFoundRequestResponse(c, "failed to find customer", err)
		return
	}

	c.JSON(200, customer)
}
