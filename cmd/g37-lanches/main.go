package main

import (
	"g37-lanchonete/configs"
	"g37-lanchonete/internal/application"
	"g37-lanchonete/internal/core/services"
	httpDriver "g37-lanchonete/internal/infra/drivers/http"
	paymentDriver "g37-lanchonete/internal/infra/drivers/payment"
	sqlDriver "g37-lanchonete/internal/infra/drivers/sql"
	"g37-lanchonete/internal/infra/gateways"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.NewConfig()
	appConfig, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	httpClient := httpDriver.NewHttpClient()
	postgresSQLClient := createPostgresSQLClient(appConfig)

	paymentBroker := paymentDriver.NewMercadoPagoBroker(httpClient, appConfig.PaymentBrokerURL)

	customerRepositoryGateway := gateways.NewCustomerRepositoryGateway(postgresSQLClient)
	productRepositoryGateway := gateways.NewProductRepositoryGateway(postgresSQLClient)
	orderRepositoryGateway := gateways.NewOrderRepositoryGateway(postgresSQLClient)

	customerService := services.NewCustomerService(customerRepositoryGateway)
	productService := services.NewProductService(productRepositoryGateway)
	paymentService := services.NewPaymentService(appConfig.NotificationURL, appConfig.SponsorId, paymentBroker)
	orderService := services.NewOrderService(customerService, paymentService, productService, orderRepositoryGateway)

	customerHandler := application.NewCustomerHandler(customerService)
	productHandler := application.NewProductHandler(productService)
	orderHandler := application.NewOrderHandler(orderService)

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/customers", customerHandler.GetCustomers)
		v1.POST("/customers", customerHandler.SaveCustomer)

		v1.GET("/products", productHandler.GetProducts)
		v1.POST("/products", productHandler.CreateProducts)
		v1.PUT("/products/:id", productHandler.UpdateProduct)
		v1.DELETE("/products/:id", productHandler.DeleteProduct)

		v1.GET("/orders", orderHandler.GetAllOrders)
		v1.GET("/orders/:id/status", orderHandler.GetOrderStatus)
		v1.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)
		v1.POST("/orders", orderHandler.CreateOrder)

	}

	router.Run(":8080")
}

func createPostgresSQLClient(appConfig configs.AppConfig) sqlDriver.SQLClient {
	db, err := sqlDriver.NewPostgresSQLClient(appConfig.DatabaseUser, appConfig.DatabasePassword, appConfig.DatabaseHost, appConfig.DatabasePort, appConfig.DatabaseName)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
