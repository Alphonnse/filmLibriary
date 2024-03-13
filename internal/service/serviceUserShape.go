package service

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
)

type ServiceUserShape interface {
	CreateActorInfo(context.Context, *model.UserModel) error
}
