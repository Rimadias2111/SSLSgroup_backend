package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	Id          uuid.UUID `gorm:"primary_key;type:uuid;"`
	From        string    `gorm:"type:varchar(50);not null"`
	To          string    `gorm:"type:varchar(50);not null"`
	PuTime      time.Time `gorm:"type:timestamp;not null"`
	DeliverTime time.Time `gorm:"type:timestamp;not null"`
	LoadedMiles int64     `gorm:"type:int;not null"`
	TotalMiles  int64     `gorm:"type:int;not null"`
	Provider    string    `gorm:"type:varchar(50);not null"`
	Cost        int64     `gorm:"type:int;not null"`
	Rate        float64   `gorm:"type:decimal(10,2);not null"`
	DriverId    uuid.UUID `gorm:"type:uuid;not null"`
	Driver      Driver    `gorm:"foreignKey:DriverId"`
	EmployeeId  uuid.UUID `gorm:"type:uuid;not null"`
	Employee    uuid.UUID `gorm:"foreignKey:EmployeeId" swaggerignore:"true"`
	CargoID     string    `gorm:"type:varchar(90); not null"`
	Status      string    `gorm:"type:varchar(50); not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}

type GetAllTransReq struct {
	Page           uint64    `json:"page"`
	Limit          uint64    `json:"limit"`
	CargoID        uuid.UUID `json:"cargo_id"`
	Provider       string    `json:"provider"`
	DriverName     string    `json:"driver_name"`
	DispatcherName string    `json:"dispatcher_name"`
}

type GetAllTransResp struct {
	Transactions []Transaction
	Count        int64
}
