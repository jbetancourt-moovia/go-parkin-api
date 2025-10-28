package main

import (
	"go-api-swagger/config"
	_ "go-api-swagger/docs"
	"go-api-swagger/internal/repositories"
	"go-api-swagger/internal/servers"
	"go-api-swagger/internal/services"
)

// @title Go Swagger API
// @version 1.0
// @description API Rest con Go y Swagger usando Fiber framework y Swaggo.
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	config.LoadEnvFile()
	config.ConnectDB()

	// Repos y servicios
	customerRepo := repositories.NewCustomerRepository(config.DB)
	customerService := services.NewCustomerService(customerRepo)

	go servers.StartRestAPIServer(customerService)
	servers.StartGrpcServer(customerService)

	select {} // mantiene el main vivo
}
