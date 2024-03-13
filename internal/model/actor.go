package model

import (
	"time"

	"github.com/google/uuid"
)

type ActorModel struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Sex       string    `json:"sex"`
	Birthday  time.Time `json:"birthday"`
	OtherInfo string    `json:"otherInfo"`
}
