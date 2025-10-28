package servers

import (
	"go-api-swagger/config"
	"go-api-swagger/internal/handlers"
	"go-api-swagger/internal/middlewares"
	"go-api-swagger/internal/repositories"
	"go-api-swagger/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func StartRestAPIServer(customerService *services.CustomerService) {
	app := fiber.New()
	app.Use(middlewares.ProtectedRoutes())

	handlers.RegisterCustomerRoutes(app, customerService)

	loginRepo := repositories.NewLoginRepository(config.DB)
	loginService := services.NewLoginService(loginRepo)
	handlers.RegisterLoginRoutes(app, loginService)

	app.Get("/ping", func(c *fiber.Ctx) error { return c.SendString("pong") })
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	log.Println("Servidor REST en http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
