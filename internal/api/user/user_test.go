package user_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alphonnse/filmLibriary/internal/api/user"
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/service/mocks"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUser(t *testing.T) {
	mockService := new(mocks.ServiceUserShape)

	var mockUserModel model.UserModel
	mockUserModel.Password = faker.Password()
	mockUserModel.Name = faker.Name()

	var responseModel model.UserRequestModel
	err := faker.FakeData(&responseModel)
	assert.NoError(t, err)
	responseModel.Role_id = 2

	mockService.On("AddUser", mock.Anything, &mockUserModel).Return(&responseModel, nil)

	requestBodyJSON, err := json.Marshal(mockUserModel)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBodyJSON))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := user.NewImplementationUser(mockService)
	handler.AddUser(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	mockService := new(mocks.ServiceUserShape)

	var mockUserModel model.UserModel
	mockUserModel.Password = faker.Password()
	mockUserModel.Name = faker.Name()

	var responseModel model.UserRequestModel
	err := faker.FakeData(&responseModel)
	assert.NoError(t, err)
	responseModel.Role_id = 2

	mockService.On("GetUser", mock.Anything, &mockUserModel).Return(&responseModel, mock.Anything, nil)

	requestBodyJSON, err := json.Marshal(mockUserModel)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBodyJSON))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := user.NewImplementationUser(mockService)
	handler.GetUser(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}
