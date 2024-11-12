package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Driver struct {
	Id          uuid.UUID `gorm:"primary_key;type:uuid;"`
	Name        string    `gorm:"type:varchar(50);not null;"`
	Surname     string    `gorm:"type:varchar(50);not null;"`
	TruckNumber int       `gorm:"type:int;not null;"`
	PhoneNumber string    `gorm:"type:varchar(20);not null;"`
	Mail        string    `gorm:"type:varchar(50);not null;"`
	Birthday    time.Time `gorm:"type:date;not null;"`
	CompanyId   uuid.UUID `gorm:"type:uuid;not null;"`
	Company     Company   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type GetAllDriversResp struct {
	Drivers []Driver `json:"drivers"`
	Count   int64    `json:"count"`
}

type GetAllDriversReq struct {
	Page        uint64 `json:"page"`
	Limit       uint64 `json:"limit"`
	TruckNumber int64  `json:"truck_number"`
	Name        string `json:"name"`
}
