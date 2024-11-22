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
	Create(ctx context.Context, company *models.Driver, tx ...*gorm.DB) (string, error)
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
	GetByUsername(ctx context.Context, username string) (*models.Employee, error)
}

type Logistic interface {
	Create(ctx context.Context, update *models.Logistic, tx ...*gorm.DB) (string, error)
	Update(ctx context.Context, update *models.Logistic, tx ...*gorm.DB) error
	Delete(ctx context.Context, req models.RequestId) error
	Get(ctx context.Context, req models.RequestId) (*models.Logistic, error)
	GetAll(ctx context.Context, req models.GetAllLogisticsReq) (*models.GetAllLogisticsResp, error)
}

type Cargo interface {
	Create(ctx context.Context, cargo *models.Cargo, tx ...*gorm.DB) (string, error)
	Update(ctx context.Context, cargo *models.Cargo, tx ...*gorm.DB) error
	Delete(ctx context.Context, req models.RequestId) error
	Get(ctx context.Context, req models.RequestId) (*models.Cargo, error)
}

type Transaction interface {
	Create(ctx context.Context, transaction *models.Transaction, tx ...*gorm.DB) (string, error)
	Update(ctx context.Context, transaction *models.Transaction) error
	Delete(ctx context.Context, req models.RequestId) error
	Get(ctx context.Context, req models.RequestId) (*models.Transaction, error)
	GetAll(ctx context.Context, req models.GetAllTransReq) (*models.GetAllTransResp, error)
}

type Performance interface {
	Create(ctx context.Context, performance *models.Performance, tx ...*gorm.DB) (string, error)
	Update(ctx context.Context, performance *models.Performance) error
	Delete(ctx context.Context, req models.RequestId) error
	Get(ctx context.Context, req models.RequestId) (*models.Performance, error)
	GetAll(ctx context.Context, req models.GetAllPerformancesReq) (*models.GetAllPerformancesResp, error)
}
