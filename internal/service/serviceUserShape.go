package service

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/google/uuid"
)

type ServiceUserShape interface {
	AddUser(context.Context, *model.UserModel) (*model.UserModel, error)
	GetUserByID(context.Context, uuid.UUID) (*model.UserModel, error)
}
