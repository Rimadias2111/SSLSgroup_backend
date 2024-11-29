package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Company struct {
	Id            uuid.UUID      `gorm:"primary_key;type:uuid;" json:"id"`
	Name          string         `gorm:"type:varchar(20);not null;" json:"name"`
	Address       string         `gorm:"type:varchar(50);not null;" json:"address"`
	Number        string         `gorm:"type:varchar(20);not null;" json:"number"`
	SCAC          string         `gorm:"type:varchar(20);not null;" json:"scac"`
	StartDate     *time.Time     `gorm:"type:date;" json:"start_date"`
	DOT           int            `gorm:"type:int;not null;" json:"dot"`
	MC            int            `gorm:"type:int;not null;" json:"mc"`
	DriversNumber int            `gorm:"type:int;not null;" json:"drivers_number"`
	Drivers       []Driver       `gorm:"foreignKey:CompanyId;" swaggerignore:"true" json:"drivers"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}

type GetAllCompaniesResp struct {
	Companies []Company `json:"companies"`
	Count     int64     `json:"count"`
}

type GetAllCompaniesReq struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}
