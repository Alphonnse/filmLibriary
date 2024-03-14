package convertor

import (
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
)

func FromDatabaseActorToActor(dbActor *generated.Actor) *model.ActorModel {
	return &model.ActorModel{
		UUID:       dbActor.ID,
		Name:       dbActor.Name,
		Sex:        dbActor.Sex,
		Birthday:   dbActor.Birthday,
		OtherInfo:  dbActor.Otherinfo.String,
		Created_at: dbActor.CreatedAt,
		Updated_at: dbActor.UpdatedAt,
	}
}
