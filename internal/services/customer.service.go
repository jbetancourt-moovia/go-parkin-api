package services

import (
	"context"
	"go-api-swagger/internal/models"
	"go-api-swagger/internal/repositories"
)

type CustomerService struct {
	repo *repositories.CustomerRepository
}

func NewCustomerService(repo *repositories.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) GetAll(ctx context.Context) (*[]models.Customer, error) {
	return s.repo.GetAll(ctx)
}

func (s *CustomerService) Create(ctx context.Context, c *models.CustomerCreate) error {
	// Aquí debería ir toda la logica de negocio, validaciones, etc.
	// Que el usuario ya exista en la DB.
	// Que el email sea unico
	// Que el telefono sea unico
	// Encriptar datos sensibles
	// etc.
	// Por simplicidad, solo delegamos al repositorio
	return s.repo.Create(ctx, c)
}

func (s *CustomerService) GetByID(ctx context.Context, id int) (*models.Customer, error) {
	// Aquí debería ir toda la logica de negocio, validaciones, etc.
	// Por simplicidad, solo delegamos al repositorio

	return s.repo.GetByID(ctx, id)
}

func (s *CustomerService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
