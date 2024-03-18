package service

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/google/uuid"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name ServiceUserShape
type ServiceUserShape interface {
	AddUser(context.Context, *model.UserModel, bool) (*model.UserRequestModel, error)
	GetUserByID(context.Context, uuid.UUID) (*model.UserModel, error)
	GetUser(context.Context, *model.UserModel) (*model.UserRequestModel, string, error)
}
