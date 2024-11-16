package storage

import (
	"backend/models"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CargoRepo struct {
	db *gorm.DB
}

func NewCargoRepo(db *gorm.DB) Cargo {
	return &CargoRepo{
		db: db,
	}
}

func (s *CargoRepo) Create(ctx context.Context, cargo *models.Cargo, tx ...*gorm.DB) (string, error) {
	var (
		query = s.db
		id    = uuid.New()
	)
	cargo.Id = id
	if len(tx) > 0 && tx[0] != nil {
		query = tx[0]
	}

	if err := query.WithContext(ctx).Create(&cargo).Error; err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *CargoRepo) Update(ctx context.Context, cargo *models.Cargo, tx ...*gorm.DB) error {
	var query = s.db

	if len(tx) > 0 && tx[0] != nil {
		query = tx[0]
	}

	result := query.WithContext(ctx).Model(&cargo).Omit("Id", "EmployeeId").Updates(cargo)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (s *CargoRepo) Delete(ctx context.Context, req models.RequestId) error {
	return s.db.WithContext(ctx).Where("id = ?", req.Id).Delete(&models.Cargo{}).Error
}

func (s *CargoRepo) Get(ctx context.Context, req models.RequestId) (*models.Cargo, error) {
	var cargo *models.Cargo

	err := s.db.WithContext(ctx).Where("id = ?", req.Id).First(&cargo).Error
	if err != nil {
		return nil, err
	}

	return cargo, nil
}
