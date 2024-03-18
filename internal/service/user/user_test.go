package user_test

import (
	"testing"

	mocksAuth "github.com/Alphonnse/filmLibriary/internal/middleware/auth-jwt/mocks"
	"github.com/Alphonnse/filmLibriary/internal/model"
	mocksRepo "github.com/Alphonnse/filmLibriary/internal/repository/mocks"
	"github.com/Alphonnse/filmLibriary/internal/service/user"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestGetUesrByID(t *testing.T) {
	uuids := []uuid.UUID{uuid.New(), uuid.New(), uuid.New(), uuid.New()}
	testCases := []struct {
		name string
		uuid uuid.UUID
		expectedUser model.UserModel
	}{
		{
			name: "first",
			uuid: uuids[0],
			expectedUser: generateFakeUesrModel(t, uuids[0]),
		},
		{
			name: "second",
			uuid: uuids[1],
			expectedUser: generateFakeUesrModel(t, uuids[1]),
		},
		{
			name: "third",
			uuid: uuids[2],
			expectedUser: generateFakeUesrModel(t, uuids[2]),
		},
		{
			name: "fourth",
			uuid: uuids[3],
			expectedUser: generateFakeUesrModel(t, uuids[3]),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(mocksRepo.Respository)
			mockAuth := new(mocksAuth.AuthService)
			handler := user.NewServiceUser(mockRepo, mockAuth)

			mockRepo.On("GetUserByID", mock.Anything, testCase.uuid).Return(&testCase.expectedUser, nil)

			res, err := handler.GetUserByID(nil, testCase.uuid)
			assert.NoError(t, err)

			assert.Equal(t, res, &testCase.expectedUser)

			mockRepo.AssertExpectations(t)
		})
	}
}

func generateFakeUesrModel(t *testing.T, id uuid.UUID) model.UserModel {
	var userModel model.UserModel
	err := faker.FakeData(&userModel)
	userModel.UUID = id
	assert.NoError(t, err)
	return userModel 
}
