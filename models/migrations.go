package models

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	err := db.Migrator().DropTable(&Company{}, &Driver{}, &Logistic{}, &Employee{})
	if err != nil {
		return err
	}
	return db.AutoMigrate(
		&Company{},
		&Driver{},
		&Logistic{},
		&Employee{},
	)
}
