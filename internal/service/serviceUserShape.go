package service

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/google/uuid"
)

type ServiceUserShape interface {
	AddUser(context.Context, *model.UserModel) (*model.UserRequestModel, error)
	GetUserByID(context.Context, uuid.UUID) (*model.UserModel, error)
	GetUser(context.Context, *model.UserModel) (*model.UserRequestModel, string, error)
}
