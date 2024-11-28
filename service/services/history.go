package services

import (
	"backend/models"
	database "backend/st_database"
	"context"
)

type HistoryService struct {
	store database.IStore
}

func NewHistoryService(store database.IStore) *HistoryService {
	return &HistoryService{store: store}
}

func (s *HistoryService) GetHistory(ctx context.Context, req models.RequestId) (*models.History, error) {
	resp, err := s.store.History().Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *HistoryService) GetAll(ctx context.Context, req models.GetAllHistoryReq) (*models.GetAllHistoryResp, error) {
	resp, err := s.store.History().GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
