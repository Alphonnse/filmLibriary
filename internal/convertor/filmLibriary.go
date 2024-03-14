package convertor

import (
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/google/uuid"
)

func FromApiAddActorInfoToService(info *model.AddActorInfoRequest) *model.ActorModel {
	return &model.ActorModel{
		Name:      info.Name,
		Sex:       info.Sex,
		Birthday:  info.Birthday,
		OtherInfo: info.OtherInfo,
	}
}

func FromApiChangeActorInfoToService(info *model.ChangeActorInfoRequest) *model.ActorModel {
	return &model.ActorModel{
		UUID:      info.UUID,
		Name:      info.Name,
		Sex:       info.Sex,
		Birthday:  info.Birthday,
		OtherInfo: info.OtherInfo,
	}
}

func FromApiRmActorInfoToService(uuid uuid.UUID) *model.ActorModel {
	return &model.ActorModel{
		UUID:      uuid,
	}
}
