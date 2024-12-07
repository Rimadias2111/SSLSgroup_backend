package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type History struct {
	Id           uuid.UUID      `gorm:"primary_key;type:uuid" json:"id"`
	DriverName   string         `gorm:"size:255;not null" json:"driver_name"`
	LogisticId   uuid.UUID      `gorm:"type:uuid; not null" json:"logistic_id"`
	FromLogistic JSONBLogistic  `gorm:"type:jsonb;" json:"from_logistic"`
	ToLogistic   JSONBLogistic  `gorm:"type:jsonb;" json:"to_logistic"`
	FromCargo    *JSONBCargo    `gorm:"type:jsonb;" json:"from_cargo"`
	ToCargo      *JSONBCargo    `gorm:"type:jsonb;" json:"to_cargo"`
	EmployeeId   uuid.UUID      `gorm:"type:uuid; not null" json:"employee_id"`
	Employee     Employee       `gorm:"foreignKey:EmployeeId;references:Id" swaggerignore:"true" json:"employee"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}

type GetAllHistoryReq struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}

type GetAllHistoryResp struct {
	Histories []History `json:"histories"`
	Count     int64     `json:"count"`
}

type JSONBLogistic struct {
	Post       bool       `json:"post"`
	Status     string     `json:"status"`
	UpdateTime time.Time  `json:"update_time"`
	StTime     *time.Time `json:"st_time"`
	State      string     `json:"state"`
	Location   string     `json:"location"`
	Notion     string     `json:"notion"`
}

func (j *JSONBLogistic) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	return json.Unmarshal(bytes, j)
}

func (j *JSONBLogistic) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type JSONBCargo struct {
	Id           uuid.UUID `json:"id"`
	CargoID      string    `json:"cargo_id"`
	Provider     string    `json:"provider"`
	LoadedMiles  int64     `json:"loaded_miles"`
	FreeMiles    int64     `json:"free_miles"`
	From         string    `json:"from"`
	To           string    `json:"to"`
	Cost         int64     `json:"cost"`
	Rate         float64   `json:"rate"`
	PickUpTime   time.Time `json:"pick_up_time"`
	DeliveryTime time.Time `json:"delivery_time"`
	EmployeeId   uuid.UUID `json:"employee_id"`
}

func (j *JSONBCargo) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	return json.Unmarshal(bytes, j)
}

func (j *JSONBCargo) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
