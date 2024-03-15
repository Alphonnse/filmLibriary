package repository

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
	"github.com/google/uuid"
)

type Respository interface {
	AddActorInfo(context.Context, generated.AddActorInfoParams) (*model.ActorModel, error)
	ChangeActorInfo(context.Context, generated.ChangeActorInfoParams) (*model.ActorModel, error)
	RmActorInfo(context.Context, uuid.UUID) error

	AddUser(context.Context, generated.AddUserParams) (*model.UserModel, error)
	GetUserByID(context.Context, uuid.UUID) (*model.UserModel, error)
	GetUser(context.Context, string) (*model.UserModel, error)
}
