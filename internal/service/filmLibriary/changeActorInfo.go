package filmlibriary

import (
	"context"
	"database/sql"
	"time"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
)

func (s *serviceLibriary) ChangeActorInfo(ctx context.Context, model *model.ActorModel) (*model.ActorModel, error) {
	changedActorInfo, err := s.libriaryRepository.ChangeActorInfo(ctx, generated.ChangeActorInfoParams{
		ID:        model.UUID,
		Name:      model.Name,
		Sex:       model.Sex,
		Birthday:  model.Birthday,
		Otherinfo: sql.NullString{String: model.OtherInfo, Valid: true},
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		return nil, err
	}
	return changedActorInfo, nil
}
