package repository

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/config"
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/convertor"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
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
