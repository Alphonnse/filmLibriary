package convertor

import (
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
)

func FromDatabaseActorToActor(dbActor *generated.Actor) *model.ActorModel {
	return &model.ActorModel{
		UUID:       dbActor.ID,
		Name:       dbActor.Name,
		Sex:        dbActor.Sex,
		Birthday:   dbActor.Birthday,
		OtherInfo:  dbActor.Otherinfo.String,
		Created_at: dbActor.CreatedAt,
		Updated_at: dbActor.UpdatedAt,
	}
}

func FromDatabaseFilmToFilmResponse(dbFilm *generated.Film) *model.FilmModelResponse {
	return &model.FilmModelResponse{
		UUID:        dbFilm.ID,
		Title:       dbFilm.Title,
		Description: dbFilm.Description.String,
		ReleaseDate: dbFilm.Releasedate.String,
		Rate:        int(dbFilm.Rate.Int32),
		Created_at:  dbFilm.CreatedAt,
		Updated_at:  dbFilm.UpdatedAt,
	}
}

func FromDatabaseFilmToFilm(dbFilm *generated.Film) *model.FilmModel {
	return &model.FilmModel{
		UUID:        dbFilm.ID,
		Title:       dbFilm.Title,
		Description: dbFilm.Description.String,
		ReleaseDate: dbFilm.Releasedate.String,
		Rate:        int(dbFilm.Rate.Int32),
		Created_at:  dbFilm.CreatedAt,
		Updated_at:  dbFilm.UpdatedAt,
	}
}
