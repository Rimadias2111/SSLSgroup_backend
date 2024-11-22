package services

import (
	"backend/models"
	database "backend/st_database"
	"context"
)

type PerformanceService struct {
	store database.IStore
}

func NewPerformanceService(store database.IStore) *PerformanceService {
	return &PerformanceService{store: store}
}

func (s *PerformanceService) Create(ctx context.Context, req *models.Performance) (string, error) {
	id, err := s.store.Performance().Create(ctx, req)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *PerformanceService) Update(ctx context.Context, req *models.Performance) error {
	err := s.store.Performance().Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *PerformanceService) Delete(ctx context.Context, req models.RequestId) error {
	err := s.store.Performance().Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *PerformanceService) Get(ctx context.Context, req models.RequestId) (*models.Performance, error) {
	resp, err := s.store.Performance().Get(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *PerformanceService) GetAll(ctx context.Context, req models.GetAllPerformancesReq) (*models.GetAllPerformancesResp, error) {
	resp, err := s.store.Performance().GetAll(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
