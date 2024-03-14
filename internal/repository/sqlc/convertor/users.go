package convertor

import (
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
)

func FromDatabaseUserToUser(dbUser *generated.User) *model.UserModel {
	return &model.UserModel{
		UUID:       dbUser.ID,
		Name:       dbUser.Name,
		Password:	dbUser.Password,
		Created_at: dbUser.CreatedAt,
		Updated_at: dbUser.UpdatedAt,
	}
}
