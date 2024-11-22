package models

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&Employee{}); err != nil {
		return err
	}
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
