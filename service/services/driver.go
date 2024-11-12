package services

import (
	"backend/models"
	database "backend/st_database"
	"context"
)

type DriverService struct {
	store database.IStore
}

func NewDriverService(store database.IStore) *DriverService {
	return &DriverService{store: store}
}

func (s *DriverService) Create(ctx context.Context, driver *models.Driver) (string, error) {
	id, err := s.store.Driver().Create(ctx, driver)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *DriverService) Update(ctx context.Context, driver *models.Driver) error {
	err := s.store.Driver().Update(ctx, driver)
	if err != nil {
		return err
	}

	return nil
}

func (s *DriverService) Delete(ctx context.Context, req models.RequestId) error {
	err := s.store.Driver().Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *DriverService) Get(ctx context.Context, req models.RequestId) (*models.Driver, error) {
	driver, err := s.store.Driver().Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (s *DriverService) GetAll(ctx context.Context, req models.GetAllDriversReq) (*models.GetAllDriversResp, error) {
	resp, err := s.store.Driver().GetAll(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
