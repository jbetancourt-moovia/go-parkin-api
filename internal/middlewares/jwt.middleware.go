package middlewares

import (
	"go-api-swagger/internal/models"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ProtectedRoutes() fiber.Handler {
	return func(c *fiber.Ctx) error {
		protected := []string{
			"/customers",
			"/vehicles",
			"/transactions",
		}

		for _, route := range protected {
			if strings.HasPrefix(c.Path(), route) {
				authHeader := c.Get("Authorization")
				if authHeader == "" {
					return c.Status(401).JSON(models.BasicResponse{
						StatusCode: 401,
						Message:    "Se requiere autenticación para acceder a este recurso",
					})
				}

				tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
				claims := jwt.MapClaims{}

				token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
					return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
				})

				if err != nil || !token.Valid {
					return c.Status(401).JSON(models.BasicResponse{
						StatusCode: 401,
						Message:    "Token inválido o expirado",
					})
				}

				// Guardar claims en contexto para el handler
				c.Locals("claims", claims)
				break
			}
		}

		return c.Next()
	}
}
