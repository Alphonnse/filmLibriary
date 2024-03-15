package model

import (
	"time"

	"github.com/google/uuid"
)

type FilmModel struct {
	UUID        uuid.UUID             `json:"uuid"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	ReleaseDate string                `json:"releaseDate"`
	Rate        int                   `json:"rate"`
	Actors      []AddActorInfoRequest `json:"actors"`
	Created_at  time.Time             `json:"created_at"`
	Updated_at  time.Time             `json:"updated_at"`
}

type FilmModelResponse struct {
	UUID        uuid.UUID    `json:"uuid"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	ReleaseDate string       `json:"releaseDate"`
	Rate        int          `json:"rate"`
	Actors      []*ActorModel `json:"actors"`
	Created_at  time.Time    `json:"created_at"`
	Updated_at  time.Time    `json:"updated_at"`
}

type AddFilmInfoRequest struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	ReleaseDate string                `json:"releaseDate"`
	Rate        int                   `json:"rate"`
	Actors      []AddActorInfoRequest `json:"actors"`
}

type ChangeFilmInfoRequest struct {
	UUID        uuid.UUID                `json:"uuid"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	ReleaseDate string                   `json:"releaseDate"`
	Rate        int                      `json:"rate"`
	Actors      []ChangeActorInfoRequest `json:"actors"`
}
