package models

import "github.com/google/uuid"

type RequestId struct {
	Id uuid.UUID `json:"id"`
}
