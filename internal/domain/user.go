package domain

import "github.com/google/uuid"

type User struct {
	UUID *uuid.UUID `json:"uuid" validate:"required,uuid"`
}
