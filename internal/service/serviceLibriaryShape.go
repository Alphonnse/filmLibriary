package service

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
)

type ServiceLibriaryShape interface {
	AddActorInfo(context.Context, *model.ActorModel) (*model.ActorModel, error)
}
