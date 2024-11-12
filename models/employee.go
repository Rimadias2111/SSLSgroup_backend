package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	Id          uuid.UUID `gorm:"primary_key;type:uuid;"`
	Name        string    `gorm:"type:varchar(30); not null"`
	Surname     string    `gorm:"type:varchar(30); not null"`
	Username    string    `gorm:"type:varchar(50); unique; not null"`
	Password    string    `gorm:"type:varchar(200); not null"`
	LogoId      string    `gorm:"size:255; default: NULL;"`
	Email       string    `gorm:"type:varchar(50); unique; not null"`
	PhoneNumber string    `gorm:"type:varchar(50); not null"`
	Birthday    time.Time `gorm:"type:date; not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type GetAllEmployeesReq struct {
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
	Search string `json:"search"`
}

type GetAllEmployeesResp struct {
	Employees []Employee `json:"employees"`
	Count     int64      `json:"count"`
}
