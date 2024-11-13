package services

import (
	"backend/models"
	database "backend/st_database"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type DriverService struct {
	store database.IStore
}

func NewDriverService(store database.IStore) *DriverService {
	return &DriverService{store: store}
}

func (s *DriverService) Create(ctx context.Context, driver *models.Driver) (string, error) {
	var id string
	db := s.store.DB()
	err := db.Transaction(func(tx *gorm.DB) error {
		idA, err := s.store.Driver().Create(ctx, driver, tx)
		if err != nil {
			return err
		}
		id = idA
		driverId, err := uuid.Parse(id)
		_, err = s.store.Logistic().Create(ctx, &models.Logistic{
			DriverId:   driverId,
			UpdateTime: time.Now(),
			CargoId:    nil,
		}, tx)
		if err != nil {
			return err
		}

		return nil
	})
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
