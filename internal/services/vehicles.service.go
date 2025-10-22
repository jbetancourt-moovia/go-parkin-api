package services

import (
	"context"
	"go-api-swagger/internal/models"
	"go-api-swagger/internal/repositories"
)

type VehicleService struct {
	repo *repositories.VehicleRepository
}

func NewVehicleService(repo *repositories.VehicleRepository) *VehicleService {
	return &VehicleService{repo: repo}
}

func (s *VehicleService) GetAll(ctx context.Context) (*[]models.Vehicle, error) {
	return s.repo.GetAll(ctx)
}

func (s *VehicleService) Create(ctx context.Context, c *models.VehicleCreate) error {
	// Aquí debería ir toda la logica de negocio, validaciones, etc.
	// Que el usuario ya exista en la DB.
	// Que el email sea unico
	// Que el telefono sea unico
	// Encriptar datos sensibles
	// etc.
	// Por simplicidad, solo delegamos al repositorio
	return s.repo.Create(ctx, c)
}

func (s *VehicleService) GetByPlate(ctx context.Context, placa string) (*models.Vehicle, error) {
	// Aquí debería ir toda la logica de negocio, validaciones, etc.
	// Por simplicidad, solo delegamos al repositorio

	return s.repo.GetByPlate(ctx, placa)
}

func (s *VehicleService) Delete(ctx context.Context, placa string) error {
	return s.repo.Delete(ctx, placa)
}
