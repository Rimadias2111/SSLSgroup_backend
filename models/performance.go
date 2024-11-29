package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Performance struct {
	Id         uuid.UUID      `gorm:"primary_key;type:uuid;" json:"id"`
	Reason     string         `gorm:"type:varchar(255);" json:"reason"`
	WhoseFault string         `gorm:"type:varchar(255);" json:"whose_fault"`
	Status     string         `gorm:"type:varchar(30);" json:"status"`
	Section    string         `gorm:"type:varchar(255);" json:"section"`
	DisputedBy string         `gorm:"type:varchar(255);" json:"disputed_by"`
	Company    string         `gorm:"type:varchar(255);" json:"company"`
	LoadId     string         `gorm:"type:varchar(255); not null;" json:"load_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}

type GetAllPerformancesReq struct {
	Page       uint64 `json:"page"`
	Limit      uint64 `json:"limit"`
	Company    string `json:"company"`
	WhoseFault string `json:"whose_fault"`
	Status     string `json:"status"`
	Section    string `json:"section"`
	DisputedBy string `json:"disputed_by"`
}

type GetAllPerformancesResp struct {
	Count        int64         `json:"count"`
	Performances []Performance `json:"performances"`
}
