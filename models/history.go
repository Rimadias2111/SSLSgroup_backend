package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type History struct {
	Id           uuid.UUID `gorm:"primary_key;type:uuid"`
	DriverName   string    `gorm:"size:255;not null"`
	LogisticId   uuid.UUID `gorm:"type:uuid; not null"`
	FromLogistic Logistic  `gorm:"type:jsonb;"`
	ToLogistic   Logistic  `gorm:"type:jsonb;"`
	FromCargo    *Cargo    `gorm:"type:jsonb;"`
	ToCargo      *Cargo    `gorm:"type:jsonb;"`
	EmployeeId   uuid.UUID `gorm:"type:uuid; not null"`
	Employee     Employee  `gorm:"foreignKey:EmployeeId;references:Id" swaggerignore:"true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}

type GetAllHistoryReq struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}

type GetAllHistoryResp struct {
	Histories []History
	Count     int64
}
