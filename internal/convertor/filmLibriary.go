package convertor

import (
	"github.com/Alphonnse/filmLibriary/internal/model"
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

func FromApiGetFilmsListToService(sortBy, orderBy string) *model.GetFilmsListService {
	return &model.GetFilmsListService{
		SortBy: sortBy,
		Order:  orderBy,
	}
}

func FromApiGetFilmsListByFragmentToService(fragmentOf, fragment string) *model.GetFilmsListByFragmentService {
	return &model.GetFilmsListByFragmentService{
		FragmentOf: fragmentOf,
		Fragment:   fragment,
	}
}
