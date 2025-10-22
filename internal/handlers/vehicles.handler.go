package handlers

import (
	"go-api-swagger/internal/models"
	"go-api-swagger/internal/services"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Constructor central
func RegisterVehicleRoutes(app *fiber.App, service *services.VehicleService) {

	app.Get("/vehicles", func(c *fiber.Ctx) error {
		return GetAllVehicles(c, service)
	})
	app.Post("/vehicles", func(c *fiber.Ctx) error {
		return CreateVehicle(c, service)
	})
	app.Get("/vehicles/placa/:placa", func(c *fiber.Ctx) error {
		return GetVehicleByPlate(c, service)
	})
	app.Delete("/vehicles/placa/:placa", func(c *fiber.Ctx) error {
		return DeleteVehicle(c, service)
	})
}

// @Summary Listar vehículos
// @Tags Vehículos
// @Produce json
// @Success 200 {object} models.ListVehicleResponse
// @Router /vehicles [get]
func GetAllVehicles(c *fiber.Ctx, service *services.VehicleService) error {

	list, err := service.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(models.ListVehicleResponse{
			StatusCode: 500,
			Data:       nil,
			Message:    err.Error(),
		})
	}

	return c.Status(200).JSON(models.ListVehicleResponse{
		StatusCode: 200,
		Data:       *list,
		Message:    "Vehículos obtenidos correctamente",
	})

}

// @Summary Crear vehículo
// @Tags Vehículos
// @Accept json
// @Produce json
// @Param vehicle body models.VehicleCreate true "Vehículo a crear"
// @Success 201 {object} models.BasicResponse
// @Router /vehicles [post]
func CreateVehicle(c *fiber.Ctx, service *services.VehicleService) error {

	var vehicle models.VehicleCreate
	if err := c.BodyParser(&vehicle); err != nil {
		return c.Status(400).JSON(models.BasicResponse{
			StatusCode: 400,
			Message:    "Error en el cuerpo de la solicitud",
		})
	}

	// Validar estructura
	if err := validate.Struct(vehicle); err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Field()+" is "+e.Tag())
		}
		return c.Status(400).JSON(fiber.Map{
			"status_code": 400,
			"errors":      errors,
		})
	}

	if err := service.Create(c.Context(), &vehicle); err != nil {
		return c.Status(500).JSON(models.BasicResponse{
			StatusCode: 500,
			Message:    err.Error(),
		})
	}

	return c.Status(201).JSON(models.BasicResponse{
		StatusCode: 201,
		Message:    "Vehículo creado correctamente",
	})

}

// @Summary Obtener vehículo por Placa
// @Tags Vehículos
// @Produce json
// @Param placa path string true "Placa del vehículo"
// @Success 200 {object} models.VehicleResponse
// @Router /vehicles/placa/{placa} [get]
func GetVehicleByPlate(c *fiber.Ctx, service *services.VehicleService) error {
	placa := c.Params("placa")
	if placa == "" {
		return c.Status(400).JSON(models.BasicResponse{
			StatusCode: 400,
			Message:    "Placa inválida",
		})
	}

	vehicle, err := service.GetByPlate(c.Context(), placa)
	if err != nil {
		return c.Status(404).JSON(models.VehicleResponse{
			StatusCode: 404,
			Data:       nil,
			Message:    "Vehículo no encontrado",
		})
	}

	return c.Status(200).JSON(models.VehicleResponse{
		StatusCode: 200,
		Data:       vehicle,
		Message:    "Vehículo encontrado correctamente",
	})
}

// @Summary Eliminar vehículo
// @Tags Vehículos
// @Produce json
// @Param placa path string true "Placa del vehículo"
// @Success 200 {object} models.BasicResponse
// @Router /vehicles/placa/{placa} [delete]
func DeleteVehicle(c *fiber.Ctx, service *services.VehicleService) error {

	placa := c.Params("placa")
	if placa == "" {
		return c.Status(400).JSON(models.BasicResponse{
			StatusCode: 400,
			Message:    "Placa inválida",
		})
	}

	if err := service.Delete(c.Context(), placa); err != nil {

		if strings.Contains(err.Error(), "no encontrado") {
			return c.Status(404).JSON(models.BasicResponse{
				StatusCode: 404,
				Message:    "Vehículo no encontrado",
			})
		}

		return c.Status(500).JSON(models.BasicResponse{
			StatusCode: 500,
			Message:    err.Error(),
		})
	}

	return c.Status(200).JSON(models.BasicResponse{
		StatusCode: 200,
		Message:    "Vehículo eliminado correctamente",
	})
}
