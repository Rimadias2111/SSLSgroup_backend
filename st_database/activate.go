package database

import (
	"backend/st_database/storage"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *Store {
	return &Store{
		db:       db,
		company:  storage.NewCompanyRepo(db),
		driver:   storage.NewDriverRepo(db),
		employee: storage.NewEmployeeRepo(db),
		logistic: storage.NewLogisticRepo(db),
	}
}
