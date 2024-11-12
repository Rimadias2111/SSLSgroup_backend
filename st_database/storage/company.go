package storage

import (
	"backend/models"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) Company {
	return &CompanyRepo{db: db}
}

func (s CompanyRepo) Create(ctx context.Context, company *models.Company) (string, error) {
	id := uuid.New()
	company.Id = id
	if err := s.db.WithContext(ctx).Create(company).Error; err != nil {
		return "", err
	}
	return id.String(), nil
}

func (s CompanyRepo) Update(ctx context.Context, company *models.Company) error {
	if err := s.db.WithContext(ctx).Model(company).Omit("Id", "DriversNumber").Updates(company).Error; err != nil {
		return err
	}
	return nil
}

func (s CompanyRepo) Delete(ctx context.Context, req models.RequestId) error {
	if err := s.db.WithContext(ctx).Where("id = ?", req.Id).Delete(&models.Company{}).Error; err != nil {
		return err
	}
	return nil
}

func (s CompanyRepo) Get(ctx context.Context, req models.RequestId) (*models.Company, error) {
	var company models.Company
	if err := s.db.WithContext(ctx).Where("id = ?", req.Id).First(&company).Error; err != nil {
		return &company, err
	}
	return &company, nil
}

func (s CompanyRepo) GetAll(ctx context.Context, req models.GetAllCompaniesReq) (*models.GetAllCompaniesResp, error) {
	var (
		resp   models.GetAllCompaniesResp
		offset = (req.Page - 1) * req.Limit
		query  = s.db.WithContext(ctx).Model(models.Company{})
	)
	if err := query.Find(&resp.Companies).Offset(int(offset)).Limit(int(req.Limit)).Error; err != nil {
		return &resp, err
	}

	if err := query.Count(&resp.Count).Error; err != nil {
		return &resp, err
	}

	return &resp, nil
}
