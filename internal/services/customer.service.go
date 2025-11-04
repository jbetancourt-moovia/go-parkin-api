package services

import (
	"context"
	"errors"
	"fmt"
	"go-api-swagger/internal/models"
	"go-api-swagger/internal/repositories"
	"strings"
)

type CustomerService struct {
	repo repositories.ICustomerRepository
}

func NewCustomerService(repo repositories.ICustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) GetAll(ctx context.Context) (*[]models.Customer, error) {
	return s.repo.GetAll(ctx)
}

func (s *CustomerService) Create(ctx context.Context, c *models.CustomerCreate) error {

	// Normalizar valores
	c.Email = strings.TrimSpace(strings.ToLower(c.Email))
	c.Phone = strings.TrimSpace(c.Phone)

	// Validaciones de negocio
	if c.Email == "" {
		return errors.New("el email es obligatorio")
	}

	// Verificar unicidad
	existingByEmail, _ := s.repo.GetByEmail(ctx, c.Email)
	if existingByEmail != nil {
		return errors.New("ya existe un cliente con ese email")
	}

	if c.Username != nil {
		username := strings.TrimSpace(*c.Username)
		if username != "" {
			existingByUsername, _ := s.repo.GetByUsername(ctx, username)
			if existingByUsername != nil {
				return errors.New("ya existe un cliente con ese nombre de usuario")
			}
		}
	}

	// Crear en DB
	if err := s.repo.Create(ctx, c); err != nil {
		return err
	}

	return nil
}

func (s *CustomerService) GetByID(ctx context.Context, id int) (*models.Customer, error) {
	if id <= 0 {
		return nil, errors.New("el ID del cliente debe ser mayor que cero")
	}

	customer, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el cliente: %w", err)
	}

	if customer == nil {
		return nil, errors.New("cliente no encontrado")
	}

	return customer, nil
}

func (s *CustomerService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("el ID del cliente debe ser mayor que cero")
	}

	customer, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("error al verificar existencia del cliente: %w", err)
	}

	if customer == nil {
		return errors.New("no se puede eliminar un cliente que no existe")
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("error al eliminar el cliente: %w", err)
	}

	return nil
}
