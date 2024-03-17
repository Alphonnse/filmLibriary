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
	UUID        uuid.UUID     `json:"uuid"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	ReleaseDate string        `json:"releaseDate"`
	Rate        int           `json:"rate"`
	Actors      []*ActorModel `json:"actors"`
	Created_at  time.Time     `json:"created_at"`
	Updated_at  time.Time     `json:"updated_at"`
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

type GetFilmsListService struct {
	SortBy string `json:"sortBy"`
	Order  string `json:"order"`
}

type FilmsListModel struct {
	UUID        uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate string    `json:"releaseDate"`
	Rate        int       `json:"rate"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type GetFilmsListByFragmentService struct {
	FragmentOf string `json:"fragmentOf"`
	Fragment   string `json:"fragment"`
}
