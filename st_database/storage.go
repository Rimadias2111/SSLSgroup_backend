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
	Cargo() storage.Cargo
	Transaction() storage.Transaction
	Performance() storage.Performance
	History() storage.History
	DB() *gorm.DB
}

type Store struct {
	db          *gorm.DB
	company     storage.Company
	driver      storage.Driver
	employee    storage.Employee
	logistic    storage.Logistic
	cargo       storage.Cargo
	transaction storage.Transaction
	performance storage.Performance
	history     storage.History
}

func (s *Store) Company() storage.Company { return s.company }

func (s *Store) Driver() storage.Driver { return s.driver }

func (s *Store) Employee() storage.Employee { return s.employee }

func (s *Store) Logistic() storage.Logistic { return s.logistic }

func (s *Store) Cargo() storage.Cargo { return s.cargo }

func (s *Store) Transaction() storage.Transaction { return s.transaction }

func (s *Store) Performance() storage.Performance { return s.performance }

func (s *Store) History() storage.History { return s.history }

func (s *Store) DB() *gorm.DB { return s.db }
