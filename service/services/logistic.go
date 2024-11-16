package services

import (
	"backend/models"
	database "backend/st_database"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (s *LogisticService) UpdateWithCargo(ctx context.Context, logistic *models.Logistic, cargo *models.Cargo, create bool) (string, error) {
	var (
		db  = s.store.DB()
		id  string
		err error
	)
	transErr := db.Transaction(func(tx *gorm.DB) error {
		if create {
			id, err = s.store.Cargo().Create(ctx, cargo, tx)
			if err != nil {
				return err
			}

			cargoId, errP := uuid.Parse(id)
			if errP != nil {
				return errP
			}

			logistic.CargoId = &cargoId
		} else {
			err = s.store.Cargo().Update(ctx, cargo, tx)
			if err != nil {
				return err
			}
			id = cargo.Id.String()
		}

		err = s.store.Logistic().Update(ctx, logistic, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if transErr != nil {
		return "", transErr
	}

	return id, nil
}
