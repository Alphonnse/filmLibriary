package repository

import (
	// "context"

	"github.com/Alphonnse/filmLibriary/internal/config"
	// "github.com/Alphonnse/filmLibriary/internal/model"
)

type repository struct {
	dbConf config.DatabaseConfig
	db *database.Queries
} 

func NewRepository(conn config.DatabaseConfig) *repository{
	db := database.New(conn.Address())
	return &repository{
		db: db,
	}
}

// func (r *repository) CreateActorInfo(context.Context, *model.UserModel) error {
// 	return nil
// }
