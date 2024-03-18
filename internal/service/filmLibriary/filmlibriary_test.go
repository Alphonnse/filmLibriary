package filmlibriary_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/mocks"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
	filmlibriary "github.com/Alphonnse/filmLibriary/internal/service/filmLibriary"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetFilmsList_SortsByVariousCriteria(t *testing.T) {
	testCases := []struct {
		name          string
		sortBy        string
		orderBy       string
		sortByCode    string
		orderByCode   string
		expectedFilms []model.FilmsListModel
	}{
		{
			name:          "Sort by title descending",
			sortBy:        "title",
			orderBy:       "descending",
			sortByCode:    "Title",
			orderByCode:   "Desc",
			expectedFilms: generateFakeFilmsList(t),
		},
		{
			name:          "Sort by release date descending",
			sortBy:        "releaseDate",
			orderBy:       "descending",
			sortByCode:    "ReleaseDate",
			orderByCode:   "Desc",
			expectedFilms: generateFakeFilmsList(t),
		},
		{
			name:          "Sort by rate descending",
			sortBy:        "rate",
			orderBy:       "descending",
			sortByCode:    "Rate",
			orderByCode:   "Desc",
			expectedFilms: generateFakeFilmsList(t),
		},
		{
			name:          "Sort by title ascending",
			sortBy:        "title",
			orderBy:       "ascending",
			sortByCode:    "Title",
			orderByCode:   "Asc",
			expectedFilms: generateFakeFilmsList(t),
		},
		{
			name:          "Sort by release date ascending",
			sortBy:        "releaseDate",
			orderBy:       "ascending",
			sortByCode:    "ReleaseDate",
			orderByCode:   "Asc",
			expectedFilms: generateFakeFilmsList(t),
		},
		{
			name:          "Sort by rate ascending",
			sortBy:        "rate",
			orderBy:       "ascending",
			sortByCode:    "Rate",
			orderByCode:   "Asc",
			expectedFilms: generateFakeFilmsList(t),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(mocks.Respository)
			handler := filmlibriary.NewServiceLibriary(mockRepo)

			mockRepo.On("GetFilmsSortedByOrderedBy"+testCase.sortByCode+testCase.orderByCode, mock.Anything).Return(testCase.expectedFilms, nil)

			res, err := handler.GetFilmsList(nil, &model.GetFilmsListService{
				SortBy: testCase.sortBy,
				Order:  testCase.orderBy,
			})

			assert.NoError(t, err)
			assert.Equal(t, res, testCase.expectedFilms)

			mockRepo.AssertExpectations(t)
		})
	}
}

func generateFakeFilmsList(t *testing.T) []model.FilmsListModel {
	var filmsListResp []model.FilmsListModel
	err := faker.FakeData(&filmsListResp)
	assert.NoError(t, err)
	return filmsListResp
}

func TestGetFilmsFragment(t *testing.T) {
	testCases := []struct {
		name          string
		fragmentOf    string
		fragment      string
		call          string
		expectedFilms []model.FilmsListModel
	}{
		{
			name:          "Get list where actor name is Fr",
			fragmentOf:    "actor_name",
			fragment:      "Fr",
			call:          "ByActorNameFragment",
			expectedFilms: generateFakeFilmsList(t),
		},
		{
			name:          "Get list where actor name is Ab",
			fragmentOf:    "actor_name",
			fragment:      "Ab",
			call:          "ByActorNameFragment",
			expectedFilms: generateFakeFilmsList(t),
		},
		{
			name:          "Get list where film title is Fr",
			fragmentOf:    "film_title",
			fragment:      "Fr",
			call:          "ListByTitleFragment",
			expectedFilms: generateFakeFilmsList(t),
		},
		{
			name:          "Get list where film title is te",
			fragmentOf:    "film_title",
			fragment:      "te",
			call:          "ListByTitleFragment",
			expectedFilms: generateFakeFilmsList(t),
		},
		{
			name:          "Get list where film title is to",
			fragmentOf:    "film_title",
			fragment:      "to",
			call:          "ListByTitleFragment",
			expectedFilms: generateFakeFilmsList(t),
		},
		{
			name:          "Get list where film title is hello",
			fragmentOf:    "film_title",
			fragment:      "hello",
			call:          "ListByTitleFragment",
			expectedFilms: generateFakeFilmsList(t),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := new(mocks.Respository)
			handler := filmlibriary.NewServiceLibriary(mockRepo)

			mockRepo.On("GetFilms"+testCase.call, mock.Anything, testCase.fragment).Return(testCase.expectedFilms, nil)

			res, err := handler.GetFilmsListByFragment(nil, &model.GetFilmsListByFragmentService{
				FragmentOf: testCase.fragmentOf,
				Fragment: testCase.fragment,
			})

			assert.NoError(t, err)
			assert.Equal(t, res, testCase.expectedFilms)

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetActorsListWithFilms(t *testing.T) {
	mockRepo := new(mocks.Respository)

	var actorsAndFilm model.GetActorsAndTeirFilmsService

	err := faker.FakeData(&actorsAndFilm)
	assert.NoError(t, err)

	mockRepo.On("GetActorsListWithFilms", mock.Anything).Return(&actorsAndFilm, nil)

	handler := filmlibriary.NewServiceLibriary(mockRepo)
	res, _ := handler.GetActorsListWithFilms(nil)

	assert.Equal(t, res, &actorsAndFilm)
	mockRepo.AssertExpectations(t)
}

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
