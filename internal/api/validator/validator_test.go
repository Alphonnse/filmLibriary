package validator_test

import (
	"errors"
	"testing"

	"github.com/Alphonnse/filmLibriary/internal/api/validator"
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestValidateUserAddRequest(t *testing.T) {
	var modelr model.Register
	err := faker.FakeData(&modelr)
	assert.NoError(t, err)
	vals := validator.ValidateUserAddRequest(&modelr)
	if len(vals) > 0 {
		if assert.Error(t, errors.New(vals.Encode())) {
			assert.Equal(t, "", errors.New(vals.Encode()))
		}
	}
}

func TestValidateUserGetRequest(t *testing.T) {
	var modell model.Login
	err := faker.FakeData(&modell)
	assert.NoError(t, err)
	vals := validator.ValidateUserGetRequest(&modell)
	if len(vals) > 0 {
		if assert.Error(t, errors.New(vals.Encode())) {
			assert.Equal(t, "", errors.New(vals.Encode()))
		}
	}
}

func TestValidateActorAddRequest(t *testing.T) {
	var modelt model.AddActorInfoRequest
	for i := 0; i < 100; i++ {
		err := faker.FakeData(&modelt)
		assert.NoError(t, err)
		modelt.Birthday = faker.Date()
		vals := validator.ValidateActorAddRequest(&modelt)
		if len(vals) > 0 {
			if assert.Error(t, errors.New(vals.Encode())) {
				assert.Equal(t, "", errors.New(vals.Encode()))
			}
		}
	}
}

func TestValidateAddFilmInfoRequest(t *testing.T) {
	var modeltt model.AddActorInfoRequest
	var modelt model.AddFilmInfoRequest
	err := faker.FakeData(&modelt)
	assert.NoError(t, err)
	modelt.ReleaseDate = faker.Date()
	modelt.Rate = 3
	err = faker.FakeData(&modeltt)
	assert.NoError(t, err)
	modeltt.Birthday = faker.Date()
	modelt.Actors = []model.AddActorInfoRequest{modeltt}

	vals := validator.ValidateAddFilmInfoRequest(&modelt)
	if len(vals) > 0 {
		if assert.Error(t, errors.New(vals.Encode())) {
			assert.Equal(t, "", errors.New(vals.Encode()))
		}
	}
}
