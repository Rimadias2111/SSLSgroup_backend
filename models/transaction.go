package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	Id           uuid.UUID      `gorm:"primary_key;type:uuid;" json:"id"`
	From         string         `gorm:"type:varchar(50);not null" json:"from"`
	To           string         `gorm:"type:varchar(50);not null" json:"to"`
	PuTime       time.Time      `gorm:"type:timestamp;not null" json:"pu_time"`
	DeliveryTime time.Time      `gorm:"type:timestamp;not null" json:"delivery_time"`
	LoadedMiles  int64          `gorm:"type:int;not null" json:"loaded_miles"`
	TotalMiles   int64          `gorm:"type:int;not null" json:"total_miles"`
	Provider     string         `gorm:"type:varchar(50);not null" json:"provider"`
	Cost         int64          `gorm:"type:int;not null" json:"cost"`
	Rate         float64        `gorm:"type:decimal(10,2);not null" json:"rate"`
	DriverId     uuid.UUID      `gorm:"type:uuid;not null" json:"driver_id"`
	Driver       Driver         `gorm:"foreignKey:DriverId" json:"driver"`
	EmployeeId   uuid.UUID      `gorm:"type:uuid;not null" json:"employee_id"`
	Employee     Employee       `gorm:"foreignKey:EmployeeId" swaggerignore:"true" json:"employee"`
	CargoID      string         `gorm:"type:varchar(90); not null" json:"cargo_id"`
	Success      bool           `gorm:"not null" json:"success"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}

type GetAllTransReq struct {
	Page           uint64 `json:"page"`
	Limit          uint64 `json:"limit"`
	CargoID        string `json:"cargo_id"`
	Provider       string `json:"provider"`
	DriverName     string `json:"driver_name"`
	DispatcherName string `json:"dispatcher_name"`
	Success        string `json:"success"`
}

type GetAllTransResp struct {
	Transactions []Transaction `json:"transactions"`
	Count        int64         `json:"count"`
}
