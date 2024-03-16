package filmlibriary

import (
	"context"
	"database/sql"
	"time"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
	"github.com/google/uuid"
)

func (s *serviceLibriary) AddActorInfo(ctx context.Context, info *model.ActorModel) (*model.ActorModel, error) {
	bdy, _ := time.Parse("2006-01-02", info.Birthday)
	response, err := s.libriaryRepository.AddActorInfo(ctx, generated.AddActorInfoParams{
		ID:        uuid.New(),
		Name:      info.Name,
		Sex:       info.Sex,
		Birthday:  bdy,
		Otherinfo: sql.NullString{String: info.OtherInfo, Valid: true},
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
