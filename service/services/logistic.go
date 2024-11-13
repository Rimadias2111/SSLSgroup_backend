package services

import (
	"backend/models"
	database "backend/st_database"
	"context"
)

type LogisticService struct {
	store database.IStore
}

func NewLogisticService(store database.IStore) *LogisticService {
	return &LogisticService{store: store}
}

func (s *LogisticService) Create(ctx context.Context, req *models.Logistic) (string, error) {
	id, err := s.store.Logistic().Create(ctx, req)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *LogisticService) Update(ctx context.Context, req *models.Logistic) error {
	err := s.store.Logistic().Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *LogisticService) Delete(ctx context.Context, req models.RequestId) error {
	err := s.store.Logistic().Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *LogisticService) Get(ctx context.Context, req models.RequestId) (*models.Logistic, error) {
	resp, err := s.store.Logistic().Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LogisticService) GetAll(ctx context.Context, req models.GetAllLogisticsReq) (*models.GetAllLogisticsResp, error) {
	resp, err := s.store.Logistic().GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
