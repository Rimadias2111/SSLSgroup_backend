package storage

import (
	"backend/models"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PerformanceRepo struct {
	db *gorm.DB
}

func NewPerformanceRepo(db *gorm.DB) Performance {
	return &PerformanceRepo{
		db: db,
	}
}

func (s *PerformanceRepo) Create(ctx context.Context, performance *models.Performance, tx ...*gorm.DB) (string, error) {
	var (
		id    = uuid.New()
		query = s.db
	)
	performance.Id = id
	if len(tx) > 0 && tx[0] != nil {
		query = tx[0]
	}

	err := query.WithContext(ctx).Create(&performance).Error
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *PerformanceRepo) Update(ctx context.Context, performance *models.Performance) error {
	return s.db.WithContext(ctx).Model(performance).
		Omit("Id", "DisputedBy").Updates(performance).Error
}

func (s *PerformanceRepo) Delete(ctx context.Context, req models.RequestId) error {
	return s.db.WithContext(ctx).Where("id = ?", req).Delete(&models.Performance{}).Error
}

func (s *PerformanceRepo) Get(ctx context.Context, req models.RequestId) (*models.Performance, error) {
	var performance models.Performance
	err := s.db.WithContext(ctx).Where("id = ?", req).First(&performance).Error
	if err != nil {
		return nil, err
	}

	return &performance, nil
}

func (s *PerformanceRepo) GetAll(ctx context.Context, req models.GetAllPerformancesReq) (*models.GetAllPerformancesResp, error) {
	var (
		resp   models.GetAllPerformancesResp
		offset = (req.Page - 1) * req.Limit
		query  = s.db.WithContext(ctx).Model(&models.Performance{})
	)

	if req.Company != "" {
		query = query.Where("company = ?", req.Company)
	}

	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	if req.Section != "" {
		query = query.Where("section ILIKE ?", "%"+req.Section+"%")
	}

	if req.WhoseFault != "" {
		query = query.Where("whose_fault = ?", req.WhoseFault)
	}

	err := query.Find(&resp.Performances).Offset(int(offset)).Limit(int(req.Limit)).Error
	if err != nil {
		return nil, err
	}

	err = query.Count(&resp.Count).Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
