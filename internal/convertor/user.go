package convertor

import "github.com/Alphonnse/filmLibriary/internal/model"

func FromApiAddUserService(info *model.Register) *model.UserModel {
	return &model.UserModel{
		Name:     info.Name,
		Password: info.Password,
	}
}

func FromApiGetUserService(info *model.Login) *model.UserModel {
	return &model.UserModel{
		Name:     info.Name,
		Password: info.Password,
	}
}

func FromServiceUserToApi(info *model.UserModel) *model.UserRequestModel {
	return &model.UserRequestModel{
		UUID:    info.UUID,
		Role_id: info.Role_id,
		Name:    info.Name,
	}
}
