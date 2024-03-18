package repository

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
	"github.com/google/uuid"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name Respository
type Respository interface {
	AddActorInfo(context.Context, generated.AddActorInfoParams) (*model.ActorModel, error)
	ChangeActorInfo(context.Context, generated.ChangeActorInfoParams) (*model.ActorModel, error)
	RmActorInfo(context.Context, uuid.UUID) error
	GetActorByID(context.Context, uuid.UUID) (*model.ActorModel, error)
	GetActorsListWithFilms(context.Context) (*model.GetActorsAndTeirFilmsService, error) 

	AddFilmInfo(context.Context, generated.AddFilmInfoParams) (*model.FilmModelResponse, error)
	InsertIntoActorToFilme(context.Context, generated.InsertIntoActorToFilmeParams) error
	ChangeFilmInfo(context.Context, generated.ChangeFilmInfoParams) (*model.FilmModelResponse, error)
	GetFilmByID(context.Context, uuid.UUID) (*model.FilmModelResponse, error)
	RmFilmInfo(context.Context, uuid.UUID) error

	GetFilmsSortedByOrderedByRateAsc(context.Context) ([]model.FilmsListModel, error)
	GetFilmsSortedByOrderedByRateDesc(context.Context) ([]model.FilmsListModel, error)
	GetFilmsSortedByOrderedByTitleAsc(context.Context) ([]model.FilmsListModel, error)
	GetFilmsSortedByOrderedByTitleDesc(context.Context) ([]model.FilmsListModel, error)
	GetFilmsSortedByOrderedByReleaseDateAsc(context.Context) ([]model.FilmsListModel, error)
	GetFilmsSortedByOrderedByReleaseDateDesc(context.Context) ([]model.FilmsListModel, error)
	GetFilmsListByTitleFragment(context.Context, string) ([]model.FilmsListModel, error) 
	GetFilmsByActorNameFragment(context.Context, string) ([]model.FilmsListModel, error)

	AddUser(context.Context, generated.AddUserParams) (*model.UserModel, error)
	GetUserByID(context.Context, uuid.UUID) (*model.UserModel, error)
	GetUser(context.Context, string) (*model.UserModel, error)
}
