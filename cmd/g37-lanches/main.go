package main

import (
	"fmt"
	"g37-lanchonete/configs"
	"g37-lanchonete/internal/application"
	"g37-lanchonete/internal/domain/models"
	"g37-lanchonete/internal/domain/services"
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

	postgresSQLClient := createPostgresSQLClient(appConfig)
	customerRepository := repositories.NewcustomerRepository(postgresSQLClient)

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
