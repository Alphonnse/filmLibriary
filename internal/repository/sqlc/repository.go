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

func (r *repository) GetActorByID(ctx context.Context, id uuid.UUID) (*model.ActorModel, error) {
	resp, err := r.db.GetActorByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return convertor.FromDatabaseActorToActor(&resp), nil
}

func (r *repository) AddFilmInfo(ctx context.Context, arg generated.AddFilmInfoParams) (*model.FilmModelResponse, error) {
	resp, err := r.db.AddFilmInfo(ctx, arg)
	if err != nil {
		return nil, err
	}
	return convertor.FromDatabaseFilmToFilmResponse(&resp), nil
}

func (r *repository) ChangeFilmInfo(ctx context.Context, arg generated.ChangeFilmInfoParams) (*model.FilmModelResponse, error) {
	resp, err := r.db.ChangeFilmInfo(ctx, arg)
	if err != nil {
		return nil, err
	}
	return convertor.FromDatabaseFilmToFilmResponse(&resp), nil
}

func (r *repository) GetFilmByID(ctx context.Context, id uuid.UUID) (*model.FilmModelResponse, error) {
	resp, err := r.db.GetFilmByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return convertor.FromDatabaseFilmToFilmResponse(&resp), nil
}

func (r *repository) InsertIntoActorToFilme(ctx context.Context, arg generated.InsertIntoActorToFilmeParams) error {
	return r.db.InsertIntoActorToFilme(ctx, arg)
}

func (r *repository) RmFilmInfo(ctx context.Context, id uuid.UUID) error {
	return r.db.RmFilmInfo(ctx, id)
}

// func (r *repository) GetActorsByFilmId(ctx context.Context, filmID uuid.UUID) ([]generated.ActorsFilm, error) {
// 	resp, err := r.db.GetActorsByFilmId(ctx, filmID)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// mb TODO
// 	return resp, nil
// }

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

func (r *repository) GetUser(ctx context.Context, arg string) (*model.UserModel, error) {
	user, err := r.db.GetUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	return convertor.FromDatabaseUserToUser(&user), nil
}

