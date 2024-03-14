package repository

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/config"
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/convertor"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
	"github.com/google/uuid"
)

type repository struct {
	dbConf config.DatabaseConfig
	db     *generated.Queries
}

func NewRepository(conn config.DatabaseConfig) *repository {
	db := generated.New(conn.Address())
	return &repository{
		db: db,
	}
}

func (r *repository) AddActorInfo(ctx context.Context, arg generated.AddActorInfoParams) (*model.ActorModel, error) {
	resp, err := r.db.AddActorInfo(ctx, arg)
	if err != nil {
		return nil, err
	}
	return convertor.FromDatabaseActorToActor(&resp), nil
}

func (r *repository) ChangeActorInfo(ctx context.Context, arg generated.ChangeActorInfoParams) (*model.ActorModel, error) {
	resp, err := r.db.ChangeActorInfo(ctx, arg)
	if err != nil {
		return nil, err
	}

	return convertor.FromDatabaseActorToActor(&resp), nil
}

func (r *repository) RmActorInfo(ctx context.Context, id uuid.UUID) error {
	return r.db.RmActorInfo(ctx, id)
}


func (r *repository) AddUser(ctx context.Context, arg generated.AddUserParams) (*model.UserModel, error) {
	resp, err := r.db.AddUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	return convertor.FromDatabaseUserToUser(&resp), nil
}

func (r *repository) GetUserByID(ctx context.Context, arg uuid.UUID) (*model.UserModel, error) {
	user, err := r.db.GetUserByID(ctx, arg)
	if err != nil {
		return nil, err
	}
	return convertor.FromDatabaseUserToUser(&user), nil
}
