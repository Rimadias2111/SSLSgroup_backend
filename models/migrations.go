package models

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Company{},
		&Driver{},
		&Logistic{},
		&Employee{},
		&Cargo{},
		&Transaction{},
		&Performance{},
	)
}
