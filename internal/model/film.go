package model

import (
	"time"

	"github.com/google/uuid"
)

type FilmModel struct {
	UUID        uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"releaseDate"`
	Rate        int       `json:"rate"`
}
