package models

import "github.com/google/uuid"

type RequestId struct {
	Id uuid.UUID `json:"id"`
}

type AuthReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
