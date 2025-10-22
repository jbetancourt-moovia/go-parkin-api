package handlers

import (
	"go-api-swagger/internal/models"
	"go-api-swagger/internal/services"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Constructor central
func RegisterTransactionsRoutes(app *fiber.App, service *services.TransactionsService) {

	app.Post("/transactions", func(c *fiber.Ctx) error {
		return CreateTransaction(c, service)
	})

	app.Get("/transactions/getInfo/:id", func(c *fiber.Ctx) error {
		return GetTransactionInfoByID(c, service)
	})

	/*
		app.Get("/transactions", func(c *fiber.Ctx) error {
			return GetAllTransactions(c, service)
		})

		app.Get("/transactions/:placa", func(c *fiber.Ctx) error {
			return GetTransactionsByPlate(c, service)
		})



		app.Put("/transactions/setInfo/:id", func(c *fiber.Ctx) error {
			return UpdateTransaction(c, service)
		})

	*/
}

// @Summary Crear una nueva transacción
// @Tags Transacciones
// @Accept json
// @Produce json
// @Param StartTiming body models.StartTiming true "Transacción a crear"
// @Success 201 {object} models.BasicResponse
// @Router /transactions [post]
func CreateTransaction(c *fiber.Ctx, service *services.TransactionsService) error {

	var transaction models.StartTiming
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(400).JSON(models.BasicResponse{
			StatusCode: 400,
			Message:    "Error en el cuerpo de la solicitud",
		})
	}

	// Validar estructura
	if err := validate.Struct(transaction); err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Field()+" is "+e.Tag())
		}
		return c.Status(400).JSON(fiber.Map{
			"status_code": 400,
			"errors":      errors,
		})
	}

	if err := service.Create(c.Context(), &transaction); err != nil {
		return c.Status(500).JSON(models.BasicResponse{
			StatusCode: 500,
			Message:    err.Error(),
		})
	}

	return c.Status(201).JSON(models.BasicResponse{
		StatusCode: 201,
		Message:    "Transacción creada correctamente",
	})
}

// @Summary Obtener información de una transacción por ID
// @Tags Transacciones
// @Produce json
// @Param id path int true "ID de la transacción"
// @Success 200 {object} models.PaymentInfo
// @Router /transactions/getInfo/{id} [get]
func GetTransactionInfoByID(c *fiber.Ctx, service *services.TransactionsService) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(models.BasicResponse{
			StatusCode: 400,
			Message:    "ID inválido",
		})
	}

	idInt, err := strconv.Atoi(id)
	transactionInfo, err := service.GetTransactionInfoByID(c.Context(), idInt)
	if err != nil {
		return c.Status(500).JSON(models.BasicResponse{
			StatusCode: 500,
			Message:    err.Error(),
		})
	}

	return c.Status(200).JSON(transactionInfo)

}
