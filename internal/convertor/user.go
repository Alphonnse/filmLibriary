package convertor

import "github.com/Alphonnse/filmLibriary/internal/model"

func FromApiAddUserService(info *model.AddUserModel) *model.UserModel {
	return &model.UserModel{
		Name:     info.Name,
		Password: info.Password,
	}
}
