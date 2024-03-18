package filmlibriary_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	filmlibriary "github.com/Alphonnse/filmLibriary/internal/api/filmLibriary"
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/service/mocks"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddActorInfo(t *testing.T) {
	mockService := new(mocks.ServiceLibriaryShape)

	var mockActorModel model.ActorModel
	err := faker.FakeData(&mockActorModel)
	assert.NoError(t, err)
	mockActorModel.Birthday = faker.Date()

	requestModel := model.ActorModel{
		Name:      mockActorModel.Name,
		Sex:       mockActorModel.Sex,
		Birthday:  mockActorModel.Birthday,
		OtherInfo: mockActorModel.OtherInfo,
	}

	mockService.On("AddActorInfo", mock.Anything, &requestModel).Return(&mockActorModel, nil)

	requestBodyJSON, err := json.Marshal(requestModel)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/actor/addActorInfo", bytes.NewBuffer(requestBodyJSON))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := filmlibriary.NewImplementationLibriary(mockService)
	handler.AddActorInfo(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestAddFilmInfo(t *testing.T) {
	mockService := new(mocks.ServiceLibriaryShape)

	var mockActorModel model.ActorModel
	err := faker.FakeData(&mockActorModel)
	assert.NoError(t, err)
	mockActorModel.Birthday = faker.Date()

	requestModel := model.AddActorInfoRequest{
		Name:      mockActorModel.Name,
		Sex:       mockActorModel.Sex,
		Birthday:  mockActorModel.Birthday,
		OtherInfo: mockActorModel.OtherInfo,
	}

	var mockFilmModel model.FilmModelResponse
	err = faker.FakeData(&mockFilmModel)
	assert.NoError(t, err)
	mockFilmModel.ReleaseDate = faker.Date()
	mockFilmModel.Actors = []*model.ActorModel{&mockActorModel}
	mockFilmModel.Rate = 3

	requestFilmModel := model.FilmModel{
		Title:       mockFilmModel.Title,
		Description: mockFilmModel.Description,
		ReleaseDate: mockFilmModel.ReleaseDate,
		Rate:        mockFilmModel.Rate,
		Actors:      []model.AddActorInfoRequest{requestModel},
	}

	mockService.On("AddFilmInfo", mock.Anything, &requestFilmModel).Return(&mockFilmModel, nil)

	requestBodyJSON, err := json.Marshal(requestFilmModel)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/film/addFilmInfo", bytes.NewBuffer(requestBodyJSON))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := filmlibriary.NewImplementationLibriary(mockService)
	handler.AddFilmInfo(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestChangeActorInfo(t *testing.T) {
	mockService := new(mocks.ServiceLibriaryShape)

	var mockActorModel model.ActorModel
	err := faker.FakeData(&mockActorModel)
	assert.NoError(t, err)
	mockActorModel.Birthday = faker.Date()

	requestModel := model.ActorModel{
		UUID:      mockActorModel.UUID,
		Name:      mockActorModel.Name,
		Sex:       mockActorModel.Sex,
		Birthday:  mockActorModel.Birthday,
		OtherInfo: mockActorModel.OtherInfo,
	}

	mockService.On("ChangeActorInfo", mock.Anything, &requestModel).Return(&mockActorModel, nil)

	requestBodyJSON, err := json.Marshal(requestModel)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/actor/changeActorInfo", bytes.NewBuffer(requestBodyJSON))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := filmlibriary.NewImplementationLibriary(mockService)
	handler.ChangeActorInfo(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestChangeFilmInfo(t *testing.T) {
	mockService := new(mocks.ServiceLibriaryShape)

	var mockActorModel model.ActorModel
	err := faker.FakeData(&mockActorModel)
	assert.NoError(t, err)
	mockActorModel.Birthday = faker.Date()

	requestModel := model.ActorModel{
		UUID:      mockActorModel.UUID,
		Name:      mockActorModel.Name,
		Sex:       mockActorModel.Sex,
		Birthday:  mockActorModel.Birthday,
		OtherInfo: mockActorModel.OtherInfo,
	}

	var mockFilmModel model.FilmModelResponse
	err = faker.FakeData(&mockFilmModel)
	assert.NoError(t, err)
	mockFilmModel.ReleaseDate = faker.Date()
	mockFilmModel.Actors = []*model.ActorModel{&requestModel}
	mockFilmModel.Rate = 3

	filmModelRecuest := model.FilmModelResponse{
		UUID:        mockFilmModel.UUID,
		Title:       mockFilmModel.Title,
		Description: mockFilmModel.Description,
		ReleaseDate: mockFilmModel.ReleaseDate,
		Rate:        mockFilmModel.Rate,
		Actors:      mockFilmModel.Actors,
	}

	mockService.On("ChangeFilmInfo", mock.Anything, &filmModelRecuest).Return(&mockFilmModel, nil)

	requestBodyJSON, err := json.Marshal(filmModelRecuest)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/film/changeFilmInfo", bytes.NewBuffer(requestBodyJSON))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := filmlibriary.NewImplementationLibriary(mockService)
	handler.ChangeFilmInfo(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetActorsInfoListWithFilms(t *testing.T) {
	mockService := new(mocks.ServiceLibriaryShape)

	var mockActorsFilmsList model.GetActorsAndTeirFilmsService
	err := faker.FakeData(&mockActorsFilmsList)
	assert.NoError(t, err)

	mockService.On("GetActorsListWithFilms", mock.Anything).Return(&mockActorsFilmsList, nil)

	req, err := http.NewRequest(http.MethodGet, "/actor/getList", bytes.NewBuffer([]byte{}))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := filmlibriary.NewImplementationLibriary(mockService)
	handler.GetActorsInfoListWithFilms(rec, req)

	assert.Equal(t, 200, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetFilmsListByFragment(t *testing.T) {
	mockService := new(mocks.ServiceLibriaryShape)

	var mockFilmsListResponse []model.FilmsListModel
	err := faker.FakeData(&mockFilmsListResponse)
	assert.NoError(t, err)

	mockGetInfo := model.GetFilmsListByFragmentService{
		FragmentOf: "film_title",
		Fragment:   "Ra",
	}

	mockService.On("GetFilmsListByFragment", mock.Anything, &mockGetInfo).Return(mockFilmsListResponse, nil)

	req, err := http.NewRequest(http.MethodGet, "/film/findFilm/film_title/Ra", bytes.NewBuffer([]byte{}))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := filmlibriary.NewImplementationLibriary(mockService)
	handler.GetFilmsListByFragment(rec, req)

	assert.Equal(t, 200, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetFilmsList(t *testing.T) {
	mockService := new(mocks.ServiceLibriaryShape)

	var mockFilmsListResponse []model.FilmsListModel
	err := faker.FakeData(&mockFilmsListResponse)
	assert.NoError(t, err)

	mockGetInfo := model.GetFilmsListService{
		SortBy: "releaseDate",
		Order:  "ascending",
	}

	mockService.On("GetFilmsList", mock.Anything, &mockGetInfo).Return(mockFilmsListResponse, nil)

	req, err := http.NewRequest(http.MethodGet, "/film/getList/releaseDate/ascending", bytes.NewBuffer([]byte{}))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := filmlibriary.NewImplementationLibriary(mockService)
	handler.GetFilmsList(rec, req)

	assert.Equal(t, 200, rec.Code)
	mockService.AssertExpectations(t)
}

func TestRmActorInfo(t *testing.T) {
	mockService := new(mocks.ServiceLibriaryShape)

	uuid := uuid.New()

	mockService.On("RmActorInfo", mock.Anything, uuid).Return(nil)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/actor/remove/%s",uuid.String()) , bytes.NewBuffer([]byte{}))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := filmlibriary.NewImplementationLibriary(mockService)
	handler.RmActorInfo(rec, req)

	assert.Equal(t, 200, rec.Code)
	mockService.AssertExpectations(t)
}

func TestRmFilmInfo(t *testing.T) {
	mockService := new(mocks.ServiceLibriaryShape)

	uuid := uuid.New()

	mockService.On("RmFilmInfo", mock.Anything, uuid).Return(nil)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/film/remove/%s",uuid.String()) , bytes.NewBuffer([]byte{}))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := filmlibriary.NewImplementationLibriary(mockService)
	handler.RmFilmInfo(rec, req)

	assert.Equal(t, 200, rec.Code)
	mockService.AssertExpectations(t)
}
