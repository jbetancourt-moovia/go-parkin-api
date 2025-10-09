package handlers

import (
	"go-api-swagger/internal/models"
	"go-api-swagger/internal/services"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// Constructor central
func RegisterCustomerRoutes(app *fiber.App, service *services.CustomerService) {

	app.Get("/customers", func(c *fiber.Ctx) error {
		return GetAllCustomers(c, service)
	})
	app.Post("/customers", func(c *fiber.Ctx) error {
		return CreateCustomer(c, service)
	})
	app.Get("/customers/:id", func(c *fiber.Ctx) error {
		return GetCustomerByID(c, service)
	})
	app.Delete("/customers/:id", func(c *fiber.Ctx) error {
		return DeleteCustomer(c, service)
	})
}

// @Summary Listar clientes
// @Tags Clientes
// @Produce json
// @Success 200 {object} models.ListCustomerResponse
// @Router /customers [get]
func GetAllCustomers(c *fiber.Ctx, service *services.CustomerService) error {

	list, err := service.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(models.ListCustomerResponse{
			StatusCode: 500,
			Data:       nil,
			Message:    err.Error(),
		})
	}

	return c.Status(200).JSON(models.ListCustomerResponse{
		StatusCode: 200,
		Data:       *list,
		Message:    "Clientes obtenidos correctamente",
	})

}

// @Summary Crear cliente
// @Tags Clientes
// @Accept json
// @Produce json
// @Param customer body models.CustomerCreate true "Cliente a crear"
// @Success 201 {object} models.BasicResponse
// @Router /customers [post]
func CreateCustomer(c *fiber.Ctx, service *services.CustomerService) error {

	var customer models.CustomerCreate
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(400).JSON(models.BasicResponse{
			StatusCode: 400,
			Message:    "Error en el cuerpo de la solicitud",
		})
	}

	// Validar estructura
	if err := validate.Struct(customer); err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Field()+" is "+e.Tag())
		}
		return c.Status(400).JSON(fiber.Map{
			"status_code": 400,
			"errors":      errors,
		})
	}

	if err := service.Create(c.Context(), &customer); err != nil {
		return c.Status(500).JSON(models.BasicResponse{
			StatusCode: 500,
			Message:    err.Error(),
		})
	}

	return c.Status(201).JSON(models.BasicResponse{
		StatusCode: 201,
		Message:    "Cliente creado correctamente",
	})

}

// @Summary Obtener cliente por ID
// @Tags Clientes
// @Produce json
// @Param id path int true "ID del cliente"
// @Success 200 {object} models.CustomerResponse
// @Router /customers/{id} [get]
func GetCustomerByID(c *fiber.Ctx, service *services.CustomerService) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(models.BasicResponse{
			StatusCode: 400,
			Message:    "ID inválido",
		})
	}

	customer, err := service.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(models.CustomerResponse{
			StatusCode: 404,
			Data:       nil,
			Message:    "Cliente no encontrado",
		})
	}

	return c.Status(200).JSON(models.CustomerResponse{
		StatusCode: 200,
		Data:       customer,
		Message:    "Cliente encontrado correctamente",
	})
}

// @Summary Eliminar cliente
// @Tags Clientes
// @Produce json
// @Param id path int true "ID del cliente"
// @Success 200 {object} models.BasicResponse
// @Router /customers/{id} [delete]
func DeleteCustomer(c *fiber.Ctx, service *services.CustomerService) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(models.BasicResponse{
			StatusCode: 400,
			Message:    "ID inválido",
		})
	}

	if err := service.Delete(c.Context(), id); err != nil {

		if strings.Contains(err.Error(), "no encontrado") {
			return c.Status(404).JSON(models.BasicResponse{
				StatusCode: 404,
				Message:    "Cliente no encontrado",
			})
		}

		return c.Status(500).JSON(models.BasicResponse{
			StatusCode: 500,
			Message:    err.Error(),
		})
	}

	return c.Status(200).JSON(models.BasicResponse{
		StatusCode: 200,
		Message:    "Cliente eliminado correctamente",
	})
}
