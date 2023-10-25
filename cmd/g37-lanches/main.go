package main

import (
	"context"
	"fmt"
	"g37-lanchonete/configs"
	"g37-lanchonete/internal/application"
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/services"
	"g37-lanchonete/internal/infra/clients"
	"g37-lanchonete/internal/infra/repositories"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
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

	postgresSQLClient := createPostgresSQLClient(appConfig)
	queue := createQueue(appConfig)

	customerRepository := repositories.NewcustomerRepository(postgresSQLClient)
	productRepository := repositories.NewProductRepository(postgresSQLClient)
	orderRepository := repositories.NewOrderRepository(postgresSQLClient)
	paymentOrderRepository := repositories.NewPaymentOrderRepository(queue)

	httpClient := clients.NewHttpClient()
	paymentBroker := clients.NewMercadoPagoBroker(httpClient, appConfig.PaymentBrokerURL)

	customerService := services.NewCustomerService(customerRepository)
	productService := services.NewProductService(productRepository)
	paymentService := services.NewPaymentService(appConfig.NotificationURL, appConfig.SponsorId, paymentBroker, paymentOrderRepository)
	orderService := services.NewOrderService(customerService, paymentService, orderRepository)

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
		v1.PUT("/products", productHandler.UpdateProduct)
		v1.DELETE("/products", productHandler.DeleteProduct)

		v1.GET("/orders", orderHandler.GetAllOrders)
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
	db.AutoMigrate(&models.Customer{})

	return clients.NewPostgresClient(db)
}

func createQueue(appConfig configs.AppConfig) clients.Queue {
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion(appConfig.SQSRegion))
	if err != nil {
		panic(fmt.Sprintf("unable to load SDK config, %v", err))
	}

	return clients.NewSQSQueue(sqs.NewFromConfig(cfg), appConfig.SQSEndpoint)
}
