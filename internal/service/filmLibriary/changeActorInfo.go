package filmlibriary

import (
	"context"
	"database/sql"
	"time"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
)

func (s *serviceLibriary) ChangeActorInfo(ctx context.Context, model *model.ActorModel) (*model.ActorModel, error) {
	prevActorInfo, err := s.libriaryRepository.GetActorByID(ctx, model.UUID)
	model = defineActorFieldsToChange(model, prevActorInfo)

	bdy, _ := time.Parse("2006-01-02", model.Birthday)
	changedActorInfo, err := s.libriaryRepository.ChangeActorInfo(ctx, generated.ChangeActorInfoParams{
		ID:        model.UUID,
		Name:      model.Name,
		Sex:       model.Sex,
		Birthday:  bdy,
		Otherinfo: sql.NullString{String: model.OtherInfo, Valid: true},
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}

	return changedActorInfo, nil
}

func defineActorFieldsToChange(model, prevActorInfo *model.ActorModel) *model.ActorModel {
	if model.Name == "" {
		model.Name = prevActorInfo.Name
	}
	if model.Sex == "" {
		model.Sex = prevActorInfo.Sex
	}
	if model.Birthday == "" {
		model.Birthday = prevActorInfo.Birthday
	}
	if model.OtherInfo == "" {
		model.OtherInfo = prevActorInfo.OtherInfo
	}
	return model
}
