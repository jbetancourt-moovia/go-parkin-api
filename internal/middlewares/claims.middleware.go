package middlewares

import (
	"go-api-swagger/internal/helpers"
	"go-api-swagger/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RequirePermission(permission int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("claims").(jwt.MapClaims)
		if !helpers.HasPermission(claims, permission) {
			return c.Status(403).JSON(models.BasicResponse{
				Message:    "Tus permisos son insuficientes para ejecutar esta acción",
				StatusCode: 403,
			})
		}
		return c.Next()
	}
}
