package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Logistic struct {
	Id               uuid.UUID  `gorm:"primary_key;type:uuid;"`
	DriverName       string     `gorm:"type:varchar(30);not null;"`
	DriverSurname    string     `gorm:"type:varchar(30);not null;"`
	SecDriverName    *string    `gorm:"type:varchar(30);default:NULL;"`
	SecDriverSurname *string    `gorm:"type:varchar(30);default:NULL;"`
	Type             string     `gorm:"type:varchar(30);not null;"`
	Status           string     `gorm:"type:varchar(30);not null;"`
	UpdateTime       time.Time  `gorm:"type:timestamp;not null;"`
	StTime           time.Time  `gorm:"type:timestamp;"`
	State            string     `gorm:"type:varchar(90);not null;"`
	Location         string     `gorm:"type:varchar(90);not null;"`
	Emoji            string     `gorm:"type:varchar(30);not null;"`
	DriverPhone      string     `gorm:"type:varchar(30);not null;"`
	CargoId          *uuid.UUID `gorm:"type:uuid;"`
	Cargo            Cargo      `gorm:"foreignKey:CargoId"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

type GetAllLogisticsReq struct {
	Page     uint64 `json:"page"`
	Limit    uint64 `json:"limit"`
	Type     string `json:"type"`
	Location string `json:"location"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}

type GetAllLogisticsResp struct {
	Logistics []Logistic
	Count     int64
}
