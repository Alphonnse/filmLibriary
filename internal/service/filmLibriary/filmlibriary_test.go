package filmlibriary_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/mocks"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddActorInfo(t *testing.T) {
	mockRepo := new(mocks.Respository)

	actorModelInput := model.ActorModel{
		Name:      "name",
		Sex:       "sex",
		Birthday:  "1221-04-08",
		OtherInfo: "some",
	}

	bdy, err := time.Parse("2006-01-02", actorModelInput.Birthday)
	assert.NoError(t, err)
	addActorInfoParams := generated.AddActorInfoParams{
		Name:      actorModelInput.Name,
		Sex:       actorModelInput.Sex,
		Birthday:  bdy,
		Otherinfo: sql.NullString{String: actorModelInput.OtherInfo, Valid: true},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	actorModelOutput := model.ActorModel{
		Name:       addActorInfoParams.Name,
		Sex:        addActorInfoParams.Sex,
		Birthday:   addActorInfoParams.Birthday.String(),
		OtherInfo:  addActorInfoParams.Otherinfo.String,
		Created_at: time.Time{},
		Updated_at: time.Time{},
	}

	mockRepo.On("AddActorInfo", mock.Anything, addActorInfoParams).Return(&actorModelOutput, nil)
}
