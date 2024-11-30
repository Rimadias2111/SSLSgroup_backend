package storage

import (
	"backend/models"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DriverRepo struct {
	db *gorm.DB
}

func NewDriverRepo(db *gorm.DB) Driver {
	return &DriverRepo{
		db: db,
	}
}

func (s *DriverRepo) Create(ctx context.Context, driver *models.Driver, tx ...*gorm.DB) (string, error) {
	var (
		id    = uuid.New()
		query = s.db
	)
	driver.Id = id

	if len(tx) > 0 && tx[0] != nil {
		query = tx[0]
	}

	if err := query.WithContext(ctx).Create(&driver).Error; err != nil {
		return "", err
	}

	var Uquery *gorm.DB
	if len(tx) > 0 && tx[0] != nil {
		Uquery = tx[0]
	}
	if Uquery != nil {
		if err := Uquery.WithContext(ctx).Model(&models.Company{}).Where("id = ?", driver.CompanyId).UpdateColumn("drivers_number",
			gorm.Expr("drivers_number + ?", 1)).Error; err != nil {
			return "", err
		}
	}

	return id.String(), nil
}

func (s *DriverRepo) Update(ctx context.Context, driver *models.Driver) error {
	if err := s.db.WithContext(ctx).Model(&driver).Omit("Id").Updates(driver).Error; err != nil {
		return err
	}

	return nil
}

func (s *DriverRepo) Delete(ctx context.Context, req models.RequestId) error {
	if err := s.db.WithContext(ctx).Where("id = ?", req.Id).Delete(&models.Driver{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *DriverRepo) Get(ctx context.Context, req models.RequestId) (*models.Driver, error) {
	var driver models.Driver
	if err := s.db.WithContext(ctx).Where("id = ?", req.Id).Preload("Company").First(&driver).Error; err != nil {
		return nil, err
	}

	return &driver, nil
}

func (s *DriverRepo) GetAll(ctx context.Context, req models.GetAllDriversReq) (*models.GetAllDriversResp, error) {
	var (
		resp   models.GetAllDriversResp
		offset = (req.Page - 1) * req.Limit
		query  = s.db.WithContext(ctx).Model(&models.Driver{})
	)
	if req.TruckNumber != "" {
		query = query.Where("truck_number = ?", req.TruckNumber)
	}

	if req.Name != "" {
		query = query.Where("name = ?", req.Name)
	}

	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}

	if req.Position != "" {
		query = query.Where("position = ?", req.Position)
	}

	if req.CompanyId != uuid.Nil {
		query = query.Where("company_id = ?", req.CompanyId)
	}

	if err := query.Preload("Company").Find(&resp.Drivers).Offset(int(offset)).Limit(int(req.Limit)).Error; err != nil {
		return nil, err
	}

	if err := query.Count(&resp.Count).Error; err != nil {
		return nil, err
	}

	return &resp, nil
}
