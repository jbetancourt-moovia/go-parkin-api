package mocks

import (
	"context"
	"go-api-swagger/internal/models"

	"github.com/stretchr/testify/mock"
)

// Mock del repositorio
type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) GetAll(ctx context.Context) (*[]models.Customer, error) {
	args := m.Called(ctx)
	return args.Get(0).(*[]models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) GetByID(ctx context.Context, id int) (*models.Customer, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) GetByEmail(ctx context.Context, email string) (*models.Customer, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) GetByUsername(ctx context.Context, username string) (*models.Customer, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) Create(ctx context.Context, c *models.CustomerCreate) error {
	args := m.Called(ctx, c)
	return args.Error(0)
}

func (m *MockCustomerRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
