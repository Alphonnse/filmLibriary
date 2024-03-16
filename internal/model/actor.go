package model

import (
	"time"

	"github.com/google/uuid"
)

type ActorModel struct {
	UUID       uuid.UUID `json:"uuid"`
	Name       string    `json:"name"`
	Sex        string    `json:"sex"`
	Birthday   string    `json:"birthday"`
	OtherInfo  string    `json:"otherInfo"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type AddActorInfoRequest struct {
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Birthday  string `json:"birthday"`
	OtherInfo string `json:"otherInfo"`
}

type ChangeActorInfoRequest struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Sex       string    `json:"sex"`
	Birthday  string    `json:"birthday"`
	OtherInfo string    `json:"otherInfo"`
}

type RmActorInfoRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type GetActorsAndTeirFilmsService struct {
	List map[string][]string `json:"actorsList"`
}
