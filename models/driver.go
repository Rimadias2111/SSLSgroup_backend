package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Driver struct {
	Id          uuid.UUID  `gorm:"primary_key;type:uuid;"`
	Name        string     `gorm:"type:varchar(50);not null;"`
	Surname     string     `gorm:"type:varchar(50);not null;"`
	Type        string     `gorm:"type:varchar(50);not null;"`
	Position    string     `gorm:"type:varchar(50);not null;"`
	TruckNumber string     `gorm:"type:varchar; not null;"`
	PhoneNumber string     `gorm:"type:varchar(20);not null;"`
	Mail        string     `gorm:"type:varchar(50);not null;"`
	Birthday    time.Time  `gorm:"type:date;not null;"`
	StartDate   *time.Time `gorm:"type:date;"`
	CompanyId   uuid.UUID  `gorm:"type:uuid;not null;"`
	Company     Company    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" swaggerignore:"true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}

type GetAllDriversResp struct {
	Drivers []Driver `json:"drivers"`
	Count   int64    `json:"count"`
}

type GetAllDriversReq struct {
	Page        uint64    `json:"page"`
	Limit       uint64    `json:"limit"`
	TruckNumber string    `json:"truck_number"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Position    string    `json:"position"`
	CompanyId   uuid.UUID `json:"company_id"`
}

func (Driver) TableName() string {
	return "drivers"
}
