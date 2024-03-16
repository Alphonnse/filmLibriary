package convertor

import (
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/google/uuid"
)

func FromApiAddActorInfoToService(info *model.AddActorInfoRequest) *model.ActorModel {
	return &model.ActorModel{
		Name:      info.Name,
		Sex:       info.Sex,
		Birthday:  info.Birthday,
		OtherInfo: info.OtherInfo,
	}
}

func FromApiChangeActorInfoToService(info *model.ChangeActorInfoRequest) *model.ActorModel {
	return &model.ActorModel{
		UUID:      info.UUID,
		Name:      info.Name,
		Sex:       info.Sex,
		Birthday:  info.Birthday,
		OtherInfo: info.OtherInfo,
	}
}

func FromApiRmActorInfoToService(uuid uuid.UUID) *model.ActorModel {
	return &model.ActorModel{
		UUID: uuid,
	}
}

func FromApiRmFilmInfoToService(uuid uuid.UUID) *model.FilmModel {
	return &model.FilmModel{
		UUID: uuid,
	}
}

func FromApiAddFilmInfoToService(info *model.AddFilmInfoRequest) *model.FilmModel {
	return &model.FilmModel{
		Title:       info.Title,
		Description: info.Description,
		ReleaseDate: info.ReleaseDate,
		Rate:        info.Rate,
		Actors:      info.Actors,
	}
}

func FromApiChangeFilmInfoToService(info *model.ChangeFilmInfoRequest) *model.FilmModelResponse {
	var acts []*model.ActorModel
	for _, actor := range info.Actors {
		acts = append(acts, FromApiChangeActorInfoToService(&actor))
	}
	return &model.FilmModelResponse{
		UUID:        info.UUID,
		Title:       info.Title,
		Description: info.Description,
		ReleaseDate: info.ReleaseDate,
		Rate:        info.Rate,
		Actors:      acts,
	}
}

func FromApiGetFilmsListToService(info *model.GetFilmsListRequest) *model.GetFilmsListService {
	return &model.GetFilmsListService{
		SortBy: info.SortBy,
		Order:  info.Order,
	}
}

func FromApiGetFilmsListByFragmentToService(info *model.GetFilmsListByFragmentRequest) *model.GetFilmsListByFragmentService {
	return &model.GetFilmsListByFragmentService{
		FragmentOf: info.FragmentOf,
		Fragment:   info.Fragment,
	}
}
