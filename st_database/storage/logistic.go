package storage

import (
	"backend/models"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogisticRepo struct {
	db *gorm.DB
}

func NewLogisticRepo(db *gorm.DB) Logistic {
	return &LogisticRepo{
		db: db,
	}
}

func (s LogisticRepo) Create(ctx context.Context, update *models.Logistic, tx ...*gorm.DB) (string, error) {
	id := uuid.New()
	update.Id = id
	err := s.db.WithContext(ctx).Create(&update).Error
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s LogisticRepo) Update(ctx context.Context, driver *models.Logistic, tx ...*gorm.DB) error {
	var query = s.db.WithContext(ctx).Model(driver)
	err := query.
		Omit("Id", "DriverName", "DriverSurname", "DriverPhone", "Type", "CargoId").Updates(driver).Error
	if err != nil {
		return err
	}

	return nil
}

func (s LogisticRepo) Delete(ctx context.Context, req models.RequestId) error {
	err := s.db.WithContext(ctx).Model(&models.Logistic{}).Where("id = ?", req.Id).Delete(&models.Logistic{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (s LogisticRepo) Get(ctx context.Context, req models.RequestId) (*models.Logistic, error) {
	var update models.Logistic

	err := s.db.WithContext(ctx).Model(&models.Logistic{}).Where("id = ?", req.Id).First(&update).Error
	if err != nil {
		return nil, err
	}

	return &update, nil
}

func (s LogisticRepo) GetAll(ctx context.Context, req models.GetAllLogisticsReq) (*models.GetAllLogisticsResp, error) {
	var (
		resp   models.GetAllLogisticsResp
		query  = s.db.WithContext(ctx).Model(&models.Logistic{})
		offset = (req.Page - 1) * req.Limit
	)

	if req.Name != "" {
		query = query.Where("driver_name ILIKE ?", "%"+req.Name+"%")
	}

	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	if req.Location != "" {
		query = query.Where("location = ?", req.Location)
	}

	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}

	err := query.Find(&resp.Logistics).Offset(int(offset)).Limit(int(req.Limit)).Error
	if err != nil {
		return nil, err
	}

	err = query.Count(&resp.Count).Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
