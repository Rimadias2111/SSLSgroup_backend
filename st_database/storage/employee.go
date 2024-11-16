package storage

import (
	"backend/models"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmployeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) Employee {
	return &EmployeeRepo{
		db: db,
	}
}

func (s *EmployeeRepo) Create(ctx context.Context, employee *models.Employee) (string, error) {
	id := uuid.New()
	employee.Id = id

	err := s.db.WithContext(ctx).Create(&employee).Error
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *EmployeeRepo) Update(ctx context.Context, employee *models.Employee) error {
	err := s.db.WithContext(ctx).Model(employee).Omit("Id", "Password").
		Updates(employee).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeRepo) Delete(ctx context.Context, req models.RequestId) error {
	err := s.db.WithContext(ctx).Where("id = ?", req.Id).Delete(&models.Employee{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeRepo) Get(ctx context.Context, req models.RequestId) (*models.Employee, error) {
	var employee models.Employee

	err := s.db.WithContext(ctx).Where("id = ?", req.Id).First(&employee).Error
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (s *EmployeeRepo) GetAll(ctx context.Context, req models.GetAllEmployeesReq) (*models.GetAllEmployeesResp, error) {
	var (
		resp   models.GetAllEmployeesResp
		offset = (req.Page - 1) * req.Limit
		query  = s.db.WithContext(ctx).Model(&models.Employee{})
	)

	if req.Search != "" {
		query.Where("Username LIKE ?", "%"+req.Search+"%")
	}

	err := query.Find(&resp.Employees).Offset(int(offset)).Limit(int(req.Page)).Error
	if err != nil {
		return nil, err
	}

	err = query.Count(&resp.Count).Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
