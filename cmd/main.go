package main

import (
	"log"

	"go-api-swagger/config"
	_ "go-api-swagger/docs"
	"go-api-swagger/internal/handlers"
	"go-api-swagger/internal/repositories"
	"go-api-swagger/internal/services"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Go Swagger API
// @version 1.0
// @description API Rest con Go y Swagger usando Fiber framework y Swaggo.
// @BasePath /
func main() {

	config.LoadEnvFile()
	config.ConnectDB()
	app := fiber.New()

	// Registrar recursos de clientes
	customerRepo := repositories.NewCustomerRepository(config.DB)
	customerService := services.NewCustomerService(customerRepo)
	handlers.RegisterCustomerRoutes(app, customerService)

	// Registrar recursos de vehículos
	vehicleRepo := repositories.NewVehicleRepository(config.DB)
	vehicleService := services.NewVehicleService(vehicleRepo)
	handlers.RegisterVehicleRoutes(app, vehicleService)

	// Registrar recursos de transacciones
	transactionRepo := repositories.NewTransactionsRepository(config.DB)
	transactionService := services.NewTransactionsService(transactionRepo)
	handlers.RegisterTransactionsRoutes(app, transactionService)

	// Registrar recursos de login
	loginRepo := repositories.NewLoginRepository(config.DB)
	loginService := services.NewLoginService(loginRepo)
	handlers.RegisterLoginRoutes(app, loginService)

	// Ruta de prueba
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	log.Println("Servidor corriendo en http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
