package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Performance struct {
	Id         uuid.UUID `gorm:"primary_key;type:uuid;"`
	Reason     string    `gorm:"type:varchar(255);"`
	WhoseFault string    `gorm:"type:varchar(255);"`
	Status     string    `gorm:"type:varchar(30);"`
	Section    string    `gorm:"type:varchar(255);"`
	DisputedBy string    `gorm:"type:varchar(255);"`
	Company    string    `gorm:"type:varchar(255);"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}

type GetAllPerformancesReq struct {
	Page       uint64 `json:"page"`
	Limit      uint64 `json:"limit"`
	Company    string `json:"company"`
	WhoseFault string `json:"whose_fault"`
	Status     string `json:"status"`
	Section    string `json:"section"`
}

type GetAllPerformancesResp struct {
	Count        int64         `json:"count"`
	Performances []Performance `json:"performances"`
}
