package main

import (
	"fmt"
	"g37-lanchonete/configs"
	"g37-lanchonete/internal/application"
	"g37-lanchonete/internal/core/domain"
	"g37-lanchonete/internal/core/services"
	"g37-lanchonete/internal/infra/clients"
	"g37-lanchonete/internal/infra/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := configs.NewConfig()
	appConfig, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	httpClient := clients.NewHttpClient()
	postgresSQLClient := createPostgresSQLClient(appConfig)
	paymentBroker := clients.NewMercadoPagoBroker(httpClient, appConfig.PaymentBrokerURL)

	customerRepository := repositories.NewCustomerRepository(postgresSQLClient)
	productRepository := repositories.NewProductRepository(postgresSQLClient)
	orderRepository := repositories.NewOrderRepository(postgresSQLClient)

	customerService := services.NewCustomerService(customerRepository)
	productService := services.NewProductService(productRepository)
	paymentService := services.NewPaymentService(appConfig.NotificationURL, appConfig.SponsorId, paymentBroker)
	orderService := services.NewOrderService(customerService, paymentService, productService, orderRepository)

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
		v1.GET("/orders/status/:id", orderHandler.GetOrderStatus)
		v1.POST("/orders", orderHandler.CreateOrder)
	}

	router.Run(":8080")
}

func createPostgresSQLClient(appConfig configs.AppConfig) clients.SQLClient {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		appConfig.DatabaseHost,
		appConfig.DatabaseUser,
		appConfig.DatabasePassword,
		appConfig.DatabaseName,
		appConfig.DatabasePort,
		appConfig.DatabaseSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.Customer{}, &domain.Order{}, &domain.OrderItem{}, &domain.Product{})

	return clients.NewPostgresClient(db)
}
