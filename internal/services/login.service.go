package services

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"errors"
	"go-api-swagger/internal/helpers"
	"go-api-swagger/internal/models"
	"go-api-swagger/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
)

type LoginService struct {
	repo *repositories.LoginRepository
}

func NewLoginService(repo *repositories.LoginRepository) *LoginService {
	return &LoginService{repo: repo}
}

func (s *LoginService) DoLogin(ctx context.Context, loginRequest models.LoginRequest) (string, error) {

	isRegistered, err := s.repo.CheckUser(ctx, loginRequest.Username)
	if err != nil {
		return "", err
	}

	if !isRegistered {
		return "", errors.New("usuario o contraseña incorrectos")
	}

	// Autenticar frente al directorio activo
	payload := models.LoginAPIRequest{
		Username: loginRequest.Username,
		Password: loginRequest.Password,
		Group:    os.Getenv("AD_GROUP"),
	}

	// Genéricos --> T
	data, err := helpers.FetchExternalData[models.LoginAPIResponse]("POST", os.Getenv("AD_BASE_URL"), payload)

	if err != nil {
		if strings.Contains(err.Error(), "externo") {
			return "", errors.New("usuario o contraseña incorrectos")
		}
		return "", err
	}

	fmt.Printf("El token retornado es: %s\n", data.Token)

	// Obtener los claims del usuario
	permissions, err := s.repo.GetUserPermissions(ctx, loginRequest.Username)
	if err != nil {
		return "", err
	}

	// Crear claims
	claims := models.LoginClaims{
		Username:    loginRequest.Username,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Crear token firmado
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SIGNING_KEY"))

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	fmt.Println("JWT generado:", tokenString)
	return tokenString, nil
}
