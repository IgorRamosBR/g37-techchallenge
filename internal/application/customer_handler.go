package application

import (
	"g37-lanchonete/internal/domain/ports"

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

func (h CustomerHandler) GetCustomers(c *gin.Context) {
	cpf := c.Query("cpf")

	customer, err := h.customerService.GetCustomerByCPF(cpf)
	if err != nil {
		c.JSON(404, err.Error())
	}

	c.JSON(200, customer)
}
