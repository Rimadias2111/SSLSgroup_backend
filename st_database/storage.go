package database

import (
	"backend/st_database/storage"
	"gorm.io/gorm"
)

type IStore interface {
	Company() storage.Company
	Driver() storage.Driver
	Employee() storage.Employee
	Logistic() storage.Logistic
	DB() *gorm.DB
}

type Store struct {
	db       *gorm.DB
	company  storage.Company
	driver   storage.Driver
	employee storage.Employee
	logistic storage.Logistic
}

func (s *Store) Company() storage.Company { return s.company }

func (s *Store) Driver() storage.Driver { return s.driver }

func (s *Store) Employee() storage.Employee { return s.employee }

func (s *Store) Logistic() storage.Logistic { return s.logistic }

func (s *Store) DB() *gorm.DB { return s.db }
