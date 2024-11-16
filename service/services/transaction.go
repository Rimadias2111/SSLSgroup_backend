package services

import (
	"backend/models"
	database "backend/st_database"
	"context"
)

type TransactionService struct {
	store database.IStore
}

func NewTransactionService(store database.IStore) *TransactionService {
	return &TransactionService{
		store: store,
	}
}

func (s *TransactionService) Create(ctx context.Context, transaction *models.Transaction) (string, error) {
	id, err := s.store.Transaction().Create(ctx, transaction)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *TransactionService) Update(ctx context.Context, transaction *models.Transaction) error {
	return s.store.Transaction().Update(ctx, transaction)
}

func (s *TransactionService) Delete(ctx context.Context, req models.RequestId) error {
	return s.store.Transaction().Delete(ctx, req)
}

func (s *TransactionService) Get(ctx context.Context, req models.RequestId) (*models.Transaction, error) {
	resp, err := s.store.Transaction().Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *TransactionService) GetAll(ctx context.Context, req models.GetAllTransReq) (*models.GetAllTransResp, error) {
	resp, err := s.store.Transaction().GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
