package storage

import (
	"backend/models"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HistoryRepo struct {
	db *gorm.DB
}

func NewHistoryRepo(db *gorm.DB) History {
	return &HistoryRepo{
		db: db,
	}
}

func (s *HistoryRepo) Create(ctx context.Context, history *models.History, tx ...*gorm.DB) (string, error) {
	var (
		id    = uuid.New()
		query = s.db
	)
	if len(tx) > 0 && tx[0] != nil {
		query = tx[0]
	}
	history.Id = id
	err := query.WithContext(ctx).Create(&history).Error
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *HistoryRepo) Update(ctx context.Context, history *models.History) error {
	err := s.db.WithContext(ctx).Model(history).Omit("Id", "EmployeeId").Updates(history).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *HistoryRepo) Delete(ctx context.Context, req models.RequestId) error {
	err := s.db.WithContext(ctx).Model(req).Where("id = ?", req.Id).Delete(&models.History{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *HistoryRepo) Get(ctx context.Context, req models.RequestId) (*models.History, error) {
	var history models.History
	err := s.db.WithContext(ctx).Model(&models.History{}).Where("id = ?", req.Id).First(&history).Error
	if err != nil {
		return nil, err
	}

	return &history, nil
}

func (s *HistoryRepo) GetAll(ctx context.Context, req models.GetAllHistoryReq) (*models.GetAllHistoryResp, error) {
	var (
		resp   models.GetAllHistoryResp
		offset = int((req.Page - 1) * req.Limit)
		query  = s.db.WithContext(ctx).Model(&models.History{})
	)

	err := query.Find(&resp.Histories).Offset(offset).Limit(int(req.Limit)).
		Error
	if err != nil {
		return nil, err
	}

	err = query.Count(&resp.Count).Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
