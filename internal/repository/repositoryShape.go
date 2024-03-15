package repository

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
	"github.com/google/uuid"
)

type Respository interface {
	AddActorInfo(context.Context, generated.AddActorInfoParams) (*model.ActorModel, error)
	ChangeActorInfo(context.Context, generated.ChangeActorInfoParams) (*model.ActorModel, error)
	RmActorInfo(context.Context, uuid.UUID) error
	GetActorByID(context.Context, uuid.UUID) (*model.ActorModel, error) 
	// GetActorsByFilmId( context.Context, uuid.UUID) ([]generated.ActorsFilm, error) 

	AddFilmInfo(context.Context, generated.AddFilmInfoParams) (*model.FilmModelResponse, error)
	InsertIntoActorToFilme(context.Context, generated.InsertIntoActorToFilmeParams) error
	ChangeFilmInfo(context.Context, generated.ChangeFilmInfoParams) (*model.FilmModelResponse, error)
	GetFilmByID(context.Context, uuid.UUID) (*model.FilmModelResponse, error) 
	RmFilmInfo(context.Context, uuid.UUID) error

	AddUser(context.Context, generated.AddUserParams) (*model.UserModel, error)
	GetUserByID(context.Context, uuid.UUID) (*model.UserModel, error)
	GetUser(context.Context, string) (*model.UserModel, error)
}
