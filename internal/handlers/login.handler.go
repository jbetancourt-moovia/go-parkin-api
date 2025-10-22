package handlers

import (
	"go-api-swagger/internal/models"
	"go-api-swagger/internal/services"

	"github.com/gofiber/fiber/v2"
)

// Constructor central
func RegisterLoginRoutes(app *fiber.App, service *services.LoginService) {

	app.Post("/login", func(c *fiber.Ctx) error {
		return setLogin(c, service)
	})

}

// @Summary Realizar login
// @Description Permite a un usuario iniciar sesión y obtener un token JWT.
// @Tags Login
// @Accept json
// @Produce json
// @Param loginRequest body models.LoginRequest true "Datos de login"
// @Success 200 {object} models.LoginSuccessResponse
// @Failure 400 {object} models.BasicResponse
// @Failure 401 {object} models.BasicResponse
// @Router /login [post]
func setLogin(c *fiber.Ctx, service *services.LoginService) error {

	var loginRequest models.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(400).JSON(models.BasicResponse{
			StatusCode: 400,
			Message:    "Error en el cuerpo de la solicitud",
		})
	}

	jwt, err := service.DoLogin(c.Context(), loginRequest)
	if err != nil {
		return c.Status(401).JSON(models.BasicResponse{
			StatusCode: 401,
			Message:    err.Error(),
		})
	}

	return c.Status(200).JSON(models.LoginSuccessResponse{
		StatusCode: 200,
		Token:      jwt,
		Message:    "Login exitoso",
	})

}
