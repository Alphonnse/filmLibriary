package model

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	UUID       uuid.UUID `json:"uuid"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Password   string    `json:"password"`
}
