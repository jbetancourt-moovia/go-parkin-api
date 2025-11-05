package tests

import (
	"context"
	"errors"
	"go-api-swagger/internal/models"
	"go-api-swagger/internal/services"
	"go-api-swagger/test/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCustomerGetAllService(t *testing.T) {
	mockRepo := new(mocks.MockCustomerRepository)
	service := services.NewCustomerService(mockRepo)

	customers := []models.Customer{{ID: 1}, {ID: 2}}
	mockRepo.On("GetAll", mock.Anything).Return(&customers, nil)

	result, err := service.GetAll(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, &customers, result)
	mockRepo.AssertExpectations(t)

}

func TestCustomerCreateService(t *testing.T) {

	ctx := context.Background()
	mockRepo := new(mocks.MockCustomerRepository)
	service := services.NewCustomerService(mockRepo)

	t.Run("1. Cliente válido y nuevo", func(t *testing.T) {
		username := "juan"
		input := &models.CustomerCreate{Email: "test@mail.com", Username: &username, Phone: "123"}

		mockRepo.On("GetByEmail", ctx, "test@mail.com").Return((*models.Customer)(nil), nil)
		mockRepo.On("GetByUsername", ctx, "juan").Return((*models.Customer)(nil), nil)
		mockRepo.On("Create", ctx, input).Return(nil)

		err := service.Create(ctx, input)
		assert.NoError(t, err)
		mockRepo.AssertCalled(t, "Create", ctx, input)
	})

	t.Run("2. Email vacío", func(t *testing.T) {
		username := "juan"
		input := &models.CustomerCreate{Email: "", Username: &username}

		err := service.Create(ctx, input)
		assert.EqualError(t, err, "el email es obligatorio")
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("3. Email duplicado", func(t *testing.T) {
		username := "juan"
		input := &models.CustomerCreate{Email: "duplicado@mail.com", Username: &username}

		mockRepo.
			On("GetByEmail", ctx, "duplicado@mail.com").
			Return(&models.Customer{ID: 1, Email: "duplicado@mail.com"}, nil)

		err := service.Create(ctx, input)
		assert.EqualError(t, err, "ya existe un cliente con ese email")
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("4. Username duplicado", func(t *testing.T) {
		username := "ajmendez"
		input := &models.CustomerCreate{Email: "unique@mail.com", Username: &username}

		mockRepo.On("GetByEmail", ctx, "unique@mail.com").Return((*models.Customer)(nil), nil)
		mockRepo.On("GetByUsername", ctx, "ajmendez").Return(&models.Customer{ID: 1, Username: &username}, nil)

		err := service.Create(ctx, input)
		assert.EqualError(t, err, "ya existe un cliente con ese nombre de usuario")
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("5. Error al crear en DB", func(t *testing.T) {
		username := "juan"
		input := &models.CustomerCreate{Email: "nuevo@mail.com", Username: &username}

		mockRepo.On("GetByEmail", ctx, "nuevo@mail.com").Return((*models.Customer)(nil), nil)
		mockRepo.On("GetByUsername", ctx, "juan").Return((*models.Customer)(nil), nil)
		mockRepo.On("Create", ctx, input).Return(errors.New("DB error"))

		err := service.Create(ctx, input)
		mockRepo.AssertCalled(t, "Create", ctx, input)
		assert.EqualError(t, err, "DB error")
	})

	t.Run("6. Username vacío", func(t *testing.T) {
		input := &models.CustomerCreate{Email: "sinuser@mail.com"}

		mockRepo.On("GetByEmail", ctx, "sinuser@mail.com").Return((*models.Customer)(nil), nil)
		mockRepo.On("Create", ctx, input).Return(nil)

		err := service.Create(ctx, input)
		assert.NoError(t, err)
		mockRepo.AssertCalled(t, "Create", ctx, input)
	})

	t.Run("7. Normalización de email y phone", func(t *testing.T) {
		username := "juan"
		input := &models.CustomerCreate{Email: "  TEST@MAIL.COM  ", Username: &username, Phone: " 12345 "}

		mockRepo.On("GetByEmail", ctx, "test@mail.com").Return(nil, nil)
		mockRepo.On("GetByUsername", ctx, "juan").Return(nil, nil)
		mockRepo.On("Create", ctx, mock.MatchedBy(func(c *models.CustomerCreate) bool {
			return c.Email == "test@mail.com" && c.Phone == "12345"
		})).Return(nil)

		err := service.Create(ctx, input)
		assert.NoError(t, err)
	})
}

func TestCustomerGetByIDService(t *testing.T) {

	ctx := context.Background()
	mockRepo := new(mocks.MockCustomerRepository)
	service := services.NewCustomerService(mockRepo)

	t.Run("1. ID inválido", func(t *testing.T) {
		c, err := service.GetByID(ctx, 0)
		assert.Nil(t, c)
		assert.EqualError(t, err, "el ID del cliente debe ser mayor que cero")
	})

	t.Run("2. Error en el repositorio", func(t *testing.T) {
		mockRepo.On("GetByID", ctx, 1).Return((*models.Customer)(nil), errors.New("db error"))
		c, err := service.GetByID(ctx, 1)
		assert.Nil(t, c)
		assert.ErrorContains(t, err, "error al obtener el cliente")
		mockRepo.AssertCalled(t, "GetByID", ctx, 1)
	})

	t.Run("3. Cliente no encontrado", func(t *testing.T) {
		mockRepo.On("GetByID", ctx, 2).Return((*models.Customer)(nil), nil)
		c, err := service.GetByID(ctx, 2)
		assert.Nil(t, c)
		assert.EqualError(t, err, "cliente no encontrado")
	})

	t.Run("4. Cliente encontrado", func(t *testing.T) {
		expected := &models.Customer{ID: 3, Email: "user@mail.com"}
		mockRepo.On("GetByID", ctx, 3).Return(expected, nil)
		c, err := service.GetByID(ctx, 3)
		assert.NoError(t, err)
		assert.Equal(t, expected, c)
	})
}

func TestCustomerDeleteService(t *testing.T) {

	ctx := context.Background()
	mockRepo := new(mocks.MockCustomerRepository)
	service := services.NewCustomerService(mockRepo)

	t.Run("1. ID inválido", func(t *testing.T) {
		err := service.Delete(ctx, 0)
		assert.EqualError(t, err, "el ID del cliente debe ser mayor que cero")
	})

	t.Run("2. Error al verificar existencia", func(t *testing.T) {
		mockRepo.On("GetByID", ctx, 1).Return((*models.Customer)(nil), errors.New("db error"))
		err := service.Delete(ctx, 1)
		assert.ErrorContains(t, err, "error al verificar existencia del cliente")
	})

	t.Run("3. Cliente no existe", func(t *testing.T) {
		mockRepo.On("GetByID", ctx, 2).Return((*models.Customer)(nil), nil)
		err := service.Delete(ctx, 2)
		assert.EqualError(t, err, "no se puede eliminar un cliente que no existe")
	})

	t.Run("4. Error al eliminar", func(t *testing.T) {
		existing := &models.Customer{ID: 3}
		mockRepo.On("GetByID", ctx, 3).Return(existing, nil)
		mockRepo.On("Delete", ctx, 3).Return(errors.New("db error"))
		err := service.Delete(ctx, 3)
		assert.ErrorContains(t, err, "error al eliminar el cliente")
	})

	t.Run("5. Eliminación exitosa", func(t *testing.T) {
		existing := &models.Customer{ID: 4}
		mockRepo.On("GetByID", ctx, 4).Return(existing, nil)
		mockRepo.On("Delete", ctx, 4).Return(nil)
		err := service.Delete(ctx, 4)
		assert.NoError(t, err)
		mockRepo.AssertCalled(t, "Delete", ctx, 4)
	})
}
