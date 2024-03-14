package convertor

import (
	"github.com/Alphonnse/filmLibriary/internal/model"
)

func FromApiAddActorInfoToService(info *model.AddActorInfoRequest) *model.ActorModel {
	return &model.ActorModel{
		Name:      info.Name,
		Sex:       info.Sex,
		Birthday:  info.Birthday,
		OtherInfo: info.OtherInfo,
	}
}
