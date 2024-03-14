package repository

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
)

type Respository interface {
	AddActorInfo(ctx context.Context, arg generated.AddActorInfoParams) (*model.ActorModel, error)
}
