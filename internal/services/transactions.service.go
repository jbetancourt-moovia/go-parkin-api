package services

import (
	"context"
	"go-api-swagger/internal/models"
	"go-api-swagger/internal/repositories"
)

type TransactionsService struct {
	repo *repositories.TransactionsRepository
}

func NewTransactionsService(repo *repositories.TransactionsRepository) *TransactionsService {
	return &TransactionsService{repo: repo}
}

func (s *TransactionsService) Create(ctx context.Context, c *models.StartTiming) error {
	return s.repo.Create(ctx, c)
}

func (s *TransactionsService) GetTransactionInfoByID(ctx context.Context, id int) (*models.PaymentInfo, error) {
	return s.repo.GetTransactionInfoByID(ctx, id)
}
