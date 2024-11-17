package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Logistic struct {
	Id         uuid.UUID  `gorm:"primary_key;type:uuid;"`
	Post       bool       `gorm:"default:false;"`
	DriverId   uuid.UUID  `gorm:"type:uuid; unique; not null"`
	Driver     Driver     `gorm:"foreignKey:DriverId;references:Id" swaggerignore:"true"`
	Status     string     `gorm:"type:varchar(30);not null; default: 'ready'"`
	UpdateTime time.Time  `gorm:"type:timestamp;not null;"`
	StTime     *time.Time `gorm:"type:timestamp;"`
	State      string     `gorm:"type:varchar(90);not null;"`
	Location   string     `gorm:"type:varchar(90);not null;"`
	Emoji      string     `gorm:"type:varchar(30);not null; default: ''"`
	Notion     string     `gorm:"type:varchar(255);not null; default: ''"`
	CargoId    *uuid.UUID `gorm:"type:uuid;"`
	Cargo      Cargo      `gorm:"foreignKey:CargoId;references:Id" swaggerignore:"true"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}

type LogisticResponse struct {
	Id             uuid.UUID
	Post           bool
	DriverId       uuid.UUID
	Status         string
	UpdateTime     time.Time
	StTime         *time.Time
	State          string
	Location       string
	Emoji          string
	Notion         string
	CargoId        *uuid.UUID
	DriverName     string
	DriverSurname  string
	DriverType     string
	DriverPosition string
	Countdown      time.Duration
	CompanyId      uuid.UUID
	CompanyName    string
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
	Logistics []LogisticResponse
	Count     int64
}
