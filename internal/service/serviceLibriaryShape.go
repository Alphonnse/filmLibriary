package service

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
)

type ServiceLibriaryShape interface {
	AddActorInfo(context.Context, *model.ActorModel) (*model.ActorModel, error)
	ChangeActorInfo(context.Context, *model.ActorModel) (*model.ActorModel, error)
	RmActorInfo(context.Context, *model.ActorModel) error
	AddFilmInfo(context.Context, *model.FilmModel) (*model.FilmModelResponse, error)
	ChangeFilmInfo(context.Context, *model.FilmModelResponse) (*model.FilmModelResponse, error)
	RmFilmInfo(context.Context, *model.FilmModel) error
	GetFilmsList(context.Context, *model.GetFilmsListService) ([]model.FilmsListModel, error)
	GetFilmsListByFragment(context.Context, *model.GetFilmsListByFragmentService) ([]model.FilmsListModel, error)
	GetActorsListWithFilms(context.Context) (*model.GetActorsAndTeirFilmsService, error) 
}
