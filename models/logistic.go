package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Logistic struct {
	Id         uuid.UUID      `gorm:"primary_key;type:uuid;" json:"id"`
	Post       bool           `gorm:"default:false;" json:"post"`
	DriverId   uuid.UUID      `gorm:"type:uuid; unique; not null" json:"driver_id"`
	Driver     Driver         `gorm:"foreignKey:DriverId;references:Id" swaggerignore:"true" json:"driver"`
	Status     string         `gorm:"type:varchar(30);not null; default: 'READY'" json:"status"`
	UpdateTime time.Time      `gorm:"type:timestamp;not null;" json:"update_time"`
	StTime     *time.Time     `gorm:"type:timestamp;" json:"st_time"`
	State      string         `gorm:"type:varchar(90);not null;" json:"state"`
	Location   string         `gorm:"type:varchar(90);not null;" json:"location"`
	Emoji      string         `gorm:"type:varchar(30);not null; default: ''" json:"emoji"`
	Notion     string         `gorm:"type:varchar(255);not null; default: ''" json:"notion"`
	CargoId    *uuid.UUID     `gorm:"type:uuid;" json:"cargo_id"`
	Cargo      Cargo          `gorm:"foreignKey:CargoId;references:Id" swaggerignore:"true" json:"cargo"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}

type LogisticResponse struct {
	Id             uuid.UUID  `json:"id"`
	Post           bool       `json:"post"`
	DriverId       uuid.UUID  `json:"driver_id"`
	Status         string     `json:"status"`
	UpdateTime     time.Time  `json:"update_time"`
	StTime         *time.Time `json:"st_time"`
	State          string     `json:"state"`
	Location       string     `json:"location"`
	Emoji          string     `json:"emoji"`
	Notion         string     `json:"notion"`
	CargoId        *uuid.UUID `json:"cargo_id"`
	DriverName     string     `json:"driver_name"`
	DriverSurname  string     `json:"driver_surname"`
	DriverType     string     `json:"driver_type"`
	DriverPosition string     `json:"driver_position"`
	Countdown      string     `json:"countdown"`
	CompanyId      uuid.UUID  `json:"company_id"`
	CompanyName    string     `json:"company_name"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type GetAllLogisticsReq struct {
	Page     uint64 `json:"page"`
	Limit    uint64 `json:"limit"`
	Post     string `json:"post"`
	Type     string `json:"type"`
	Position string `json:"position"`
	State    string `json:"state"`
	Location string `json:"location"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}

type ByCompany struct {
	CompanyId   uuid.UUID          `json:"company_id"`
	CompanyName string             `json:"company_name"`
	Logistics   []LogisticResponse `json:"logistics"`
}

type GetAllLogisticsResp struct {
	Companies []ByCompany `json:"companies"`
	Count     int64       `json:"count"`
}

type GetOverview struct {
	Companies []struct {
		Id                uuid.UUID
		Name              string
		FreeDrivers       int64
		WillBeSoonDrivers int64
		OccupiedDrivers   int64
		NotWorking        int64
	} `json:"companies"`
}

type JSONBLogistic struct {
	Logistic Logistic
}

func (j *JSONBLogistic) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, &j.Logistic)
}

func (j *JSONBLogistic) Value() (driver.Value, error) {
	return json.Marshal(j.Logistic)
}
