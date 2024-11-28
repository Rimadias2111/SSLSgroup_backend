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
	Id           uuid.UUID `gorm:"primary_key;type:uuid;not_null"`
	CargoID      string    `gorm:"type:varchar(90); not null"`
	Provider     string    `gorm:"type:varchar(90); not null"`
	LoadedMiles  int64     `gorm:"not null"`
	FreeMiles    int64     `gorm:"not null"`
	From         string    `gorm:"type:varchar(90); not null"`
	To           string    `gorm:"type:varchar(90); not null"`
	Cost         int64     `gorm:"not null"`
	Rate         float64   `gorm:"type:decimal(10,2);not null"`
	PickUpTime   time.Time `gorm:"type:timestamp;not null"`
	DeliveryTime time.Time `gorm:"type:timestamp;not null"`
	EmployeeId   uuid.UUID `gorm:"type:uuid;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
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
