package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alphonnse/filmLibriary/internal/convertor"
	"github.com/Alphonnse/filmLibriary/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *serviceUser) GetUser(ctx context.Context, model *model.UserModel) (*model.UserRequestModel, string , error) {

	userInfo, err := s.UserRepository.GetUser(ctx, model.Name)
	if err != nil {
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(model.Password))
	if err != nil {
		return nil, "" ,errors.New(fmt.Sprintf("Invalid username or password"))
	}

	tokenString, err := s.AuthService.GenerateJWT(ctx, userInfo)
	if err != nil {
		return nil, "", err
	}

	return convertor.FromServiceUserToApi(userInfo), tokenString, nil
}
