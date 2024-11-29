package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Driver struct {
	Id          uuid.UUID      `gorm:"primary_key;type:uuid;" json:"id"`
	Name        string         `gorm:"type:varchar(50);not null;" json:"name"`
	Surname     string         `gorm:"type:varchar(50);not null;" json:"surname"`
	Type        string         `gorm:"type:varchar(50);not null;" json:"type"`
	Position    string         `gorm:"type:varchar(50);not null;" json:"position"`
	TruckNumber string         `gorm:"type:varchar; not null;" json:"truck_number"`
	PhoneNumber string         `gorm:"type:varchar(20);not null;" json:"phone_number"`
	Mail        string         `gorm:"type:varchar(50);not null;" json:"mail"`
	Birthday    time.Time      `gorm:"type:date;not null;" json:"birthday"`
	StartDate   *time.Time     `gorm:"type:date;" json:"start_date"`
	CompanyId   uuid.UUID      `gorm:"type:uuid;not null;" json:"company_id"`
	Company     Company        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" swaggerignore:"true" json:"company"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
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
