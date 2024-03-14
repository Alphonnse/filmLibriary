package user

import (
	"context"
	"time"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
	"github.com/google/uuid"
)

func (s *serviceUser) AddUser(ctx context.Context, model *model.UserModel) (*model.UserModel, error) {
	addedUserInfo, err := s.UserRepository.AddUser(ctx, generated.AddUserParams{
		ID:        uuid.New(),
		RoleID:    2,
		Name:      model.Name,
		Password:  model.Password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}
	return addedUserInfo, nil
}
