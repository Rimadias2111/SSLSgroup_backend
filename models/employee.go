package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	Id          uuid.UUID      `gorm:"primary_key;type:uuid;" json:"id"`
	Name        string         `gorm:"type:varchar(30); not null" json:"name"`
	Surname     string         `gorm:"type:varchar(30); not null" json:"surname"`
	Username    string         `gorm:"type:varchar(50); unique; not null" json:"username"`
	Position    string         `gorm:"type:varchar(30); not null" json:"position"`
	AccessLevel int64          `gorm:"type:int; not null; default:3" json:"access_level"`
	Password    string         `gorm:"type:varchar(200); not null" json:"password"`
	LogoId      string         `gorm:"size:255; default: NULL;" json:"logo_id"`
	Email       string         `gorm:"type:varchar(50); unique; not null" json:"email"`
	PhoneNumber string         `gorm:"type:varchar(50); not null" json:"phone_number"`
	Birthday    time.Time      `gorm:"type:date; not null" json:"birthday"`
	Company     string         `gorm:"type:varchar(50); default: NULL" json:"company"`
	StartDate   *time.Time     `gorm:"type:date;" json:"start_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}

type GetAllEmployeesReq struct {
	Page     uint64 `json:"page"`
	Limit    uint64 `json:"limit"`
	Search   string `json:"search"`
	Position string `json:"position"`
}

type GetAllEmployeesResp struct {
	Employees []Employee `json:"employees"`
	Count     int64      `json:"count"`
}
