package main

import (
	"g37-lanchonete/internal/application"
	"g37-lanchonete/internal/domain/services"
	"g37-lanchonete/internal/infra/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	customerRepository := repositories.NewcustomerRepository()

	customerService := services.NewCustomerService(customerRepository)

	applicationHandler := application.NewCustomerHandler(customerService)

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/customers", applicationHandler.GetCustomers)
		v1.POST("/customers")

		v1.GET("/products")
		v1.POST("/products")
		v1.PUT("/products")
		v1.DELETE("/products")
	}

	router.Run(":8080")

}
