package filmlibriary

import (
	"context"
	"database/sql"
	"time"

	"github.com/Alphonnse/filmLibriary/internal/convertor"
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
	"github.com/google/uuid"
)

func (s *serviceLibriary) AddFilmInfo(ctx context.Context, info *model.FilmModel) (*model.FilmModelResponse, error) {
	relDate, err := time.Parse("2006-01-02", info.ReleaseDate)
	if err != nil {
		return nil, err	
	}
	addedFilmInfo, err := s.libriaryRepository.AddFilmInfo(ctx, generated.AddFilmInfoParams{
		ID:          uuid.New(),
		Title:       info.Title,
		Description: sql.NullString{String: info.Description, Valid: true},
		Releasedate: sql.NullTime{Time: relDate, Valid: true},
		Rate:        sql.NullInt32{Int32: int32(info.Rate), Valid: true},
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}

	for _, actor := range info.Actors {
		addedActorModel, err := s.AddActorInfo(ctx, convertor.FromApiAddActorInfoToService(&actor))
		if err != nil {
			return addedFilmInfo, err
		}
		addedFilmInfo.Actors = append(addedFilmInfo.Actors, addedActorModel)
		s.libriaryRepository.InsertIntoActorToFilme(ctx, generated.InsertIntoActorToFilmeParams{
			FilmID:  addedFilmInfo.UUID,
			ActorID: addedActorModel.UUID,
		})
	}

	return addedFilmInfo, nil
}
