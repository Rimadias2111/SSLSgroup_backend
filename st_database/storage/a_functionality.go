package storage

import (
	"backend/models"
	"context"
	"gorm.io/gorm"
)

type Company interface {
	Create(ctx context.Context, company *models.Company) (string, error)
	Update(ctx context.Context, company *models.Company) error
	Delete(ctx context.Context, req models.RequestId) error
	Get(ctx context.Context, req models.RequestId) (*models.Company, error)
	GetAll(ctx context.Context, req models.GetAllCompaniesReq) (*models.GetAllCompaniesResp, error)
}

type Driver interface {
	Create(ctx context.Context, company *models.Driver) (string, error)
	Update(ctx context.Context, company *models.Driver) error
	Delete(ctx context.Context, req models.RequestId) error
	Get(ctx context.Context, req models.RequestId) (*models.Driver, error)
	GetAll(ctx context.Context, req models.GetAllDriversReq) (*models.GetAllDriversResp, error)
}

type Employee interface {
	Create(ctx context.Context, company *models.Employee) (string, error)
	Update(ctx context.Context, company *models.Employee) error
	Delete(ctx context.Context, req models.RequestId) error
	Get(ctx context.Context, req models.RequestId) (*models.Employee, error)
	GetAll(ctx context.Context, req models.GetAllEmployeesReq) (*models.GetAllEmployeesResp, error)
}

type Logistic interface {
	Create(ctx context.Context, company *models.Logistic, tx ...*gorm.DB) (string, error)
	Update(ctx context.Context, company *models.Logistic, tx ...*gorm.DB) error
	Delete(ctx context.Context, req models.RequestId) error
	Get(ctx context.Context, req models.RequestId) (*models.Logistic, error)
	GetAll(ctx context.Context, req models.GetAllLogisticsReq) (*models.GetAllLogisticsResp, error)
}
