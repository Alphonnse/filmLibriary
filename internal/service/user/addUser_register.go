package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Alphonnse/filmLibriary/internal/convertor"
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *serviceUser) AddUser(ctx context.Context, model *model.UserModel) (*model.UserRequestModel, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(model.Password), 10) // last argument is the coast
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error while hashing the password:", err))
	}

	addedUserInfo, err := s.UserRepository.AddUser(ctx, generated.AddUserParams{
		ID:        uuid.New(),
		RoleID:    int32(2),
		Name:      model.Name,
		Password:  string(hash),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}
	return convertor.FromServiceUserToApi(addedUserInfo), nil
}
