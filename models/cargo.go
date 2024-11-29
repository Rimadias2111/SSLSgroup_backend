package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Cargo struct {
	Id           uuid.UUID      `gorm:"primary_key;type:uuid;not_null" json:"id"`
	CargoID      string         `gorm:"type:varchar(90); not null"     json:"cargo_id"`
	Provider     string         `gorm:"type:varchar(90); not null"  json:"provider"`
	LoadedMiles  int64          `gorm:"not null" json:"loaded_miles"`
	FreeMiles    int64          `gorm:"not null" json:"free_miles"`
	From         string         `gorm:"type:varchar(90); not null" json:"from"`
	To           string         `gorm:"type:varchar(90); not null" json:"to"`
	Cost         int64          `gorm:"not null" json:"cost"`
	Rate         float64        `gorm:"type:decimal(10,2);not null" json:"rate"`
	PickUpTime   time.Time      `gorm:"type:timestamp;not null" json:"pick_up_time"`
	DeliveryTime time.Time      `gorm:"type:timestamp;not null" json:"delivery_time"`
	EmployeeId   uuid.UUID      `gorm:"type:uuid;" json:"employee_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}

type JSONBCargo struct {
	Cargo Cargo
}

func (j *JSONBCargo) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, &j.Cargo)
}

func (j *JSONBCargo) Value() (driver.Value, error) {
	return json.Marshal(j.Cargo)
}
