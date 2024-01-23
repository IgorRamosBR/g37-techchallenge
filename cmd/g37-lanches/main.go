package main

import (
	"g37-lanchonete/configs"
	"g37-lanchonete/internal/api"
	"g37-lanchonete/internal/controllers"
	"g37-lanchonete/internal/core/usecases"
	httpDriver "g37-lanchonete/internal/infra/drivers/http"
	paymentDriver "g37-lanchonete/internal/infra/drivers/payment"
	sqlDriver "g37-lanchonete/internal/infra/drivers/sql"
	"g37-lanchonete/internal/infra/gateways"
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

	customerUsecase := usecases.NewCustomerUsecase(customerRepositoryGateway)
	productUsecase := usecases.NewProductUsecase(productRepositoryGateway)
	paymentUsecase := usecases.NewPaymentUsecase(appConfig.NotificationURL, appConfig.SponsorId, paymentBroker)
	orderUsecase := usecases.NewOrderUsecase(customerUsecase, paymentUsecase, productUsecase, orderRepositoryGateway)

	customerController := controllers.NewCustomerController(customerUsecase)
	productController := controllers.NewProductController(productUsecase)
	orderController := controllers.NewOrderController(orderUsecase)

	apiParams := api.ApiParams{
		CustomerController: customerController,
		ProductController:  productController,
		OrderController:    orderController,
	}
	api := api.NewApi(apiParams)
	api.Run(":8080")
}

func createPostgresSQLClient(appConfig configs.AppConfig) sqlDriver.SQLClient {
	db, err := sqlDriver.NewPostgresSQLClient(appConfig.DatabaseUser, appConfig.DatabasePassword, appConfig.DatabaseHost, appConfig.DatabasePort, appConfig.DatabaseName)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
