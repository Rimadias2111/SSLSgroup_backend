package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Company struct {
	Id            uuid.UUID `gorm:"primary_key;type:uuid;"`
	Name          string    `gorm:"type:varchar(20);not null;"`
	Address       string    `gorm:"type:varchar(50);not null;"`
	Number        string    `gorm:"type:varchar(20);not null;"`
	SCAC          string    `gorm:"type:varchar(20);not null;"`
	DOT           int       `gorm:"type:int;not null;"`
	MC            int       `gorm:"type:int;not null;"`
	DriversNumber int       `gorm:"type:int;not null;"`
	Drivers       []Driver  `gorm:"foreignKey:CompanyId;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type GetAllCompaniesResp struct {
	Companies []Company `json:"companies"`
	Count     int64     `json:"count"`
}

type GetAllCompaniesReq struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}
