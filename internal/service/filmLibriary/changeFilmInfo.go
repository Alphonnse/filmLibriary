package filmlibriary

import (
	"context"
	"database/sql"
	"time"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
)

func (s *serviceLibriary) ChangeFilmInfo(ctx context.Context, model *model.FilmModelResponse) (*model.FilmModelResponse, error) {
	prevFilmInfo, err := s.libriaryRepository.GetFilmByID(ctx, model.UUID)
	model = defineFilmFieldsToChange(model, prevFilmInfo)

	relDate, err := time.Parse("2006-01-02", model.ReleaseDate)
	if err != nil {
		return nil, err
	}
	changedFilmInfo, err := s.libriaryRepository.ChangeFilmInfo(ctx, generated.ChangeFilmInfoParams{
		ID:          model.UUID,
		Title:       model.Title,
		Description: sql.NullString{String: model.Description, Valid: true},
		Releasedate: sql.NullTime{Time: relDate, Valid: true},
		Rate:        sql.NullInt32{Int32: int32(model.Rate), Valid: true},
		UpdatedAt:   time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}

	for _, actor := range model.Actors {
		changtedActorInfo, err := s.ChangeActorInfo(ctx, actor)
		if err != nil {
			return changedFilmInfo, err
		}
		changedFilmInfo.Actors = append(changedFilmInfo.Actors, changtedActorInfo)
	}

	return changedFilmInfo, nil
}

func defineFilmFieldsToChange(model *model.FilmModelResponse, prevActorInfo *model.FilmModelResponse) *model.FilmModelResponse {
	if model.Title == "" {
		model.Title = prevActorInfo.Title
	}
	if model.Description == "" {
		model.Description = prevActorInfo.Description
	}
	if model.ReleaseDate == "" {
		model.ReleaseDate = prevActorInfo.ReleaseDate
	}
	if model.Rate == 0 {
		model.Rate = prevActorInfo.Rate
	}
	return model
}
