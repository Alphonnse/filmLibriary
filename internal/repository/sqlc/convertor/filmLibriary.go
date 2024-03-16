package convertor

import (
	"strings"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/repository/sqlc/generated"
)

func FromDatabaseActorToActor(dbActor *generated.Actor) *model.ActorModel {
	return &model.ActorModel{
		UUID:       dbActor.ID,
		Name:       dbActor.Name,
		Sex:        dbActor.Sex,
		Birthday:   dbActor.Birthday.String(),
		OtherInfo:  dbActor.Otherinfo.String,
		Created_at: dbActor.CreatedAt,
		Updated_at: dbActor.UpdatedAt,
	}
}

func FromDatabaseFilmToFilmResponse(dbFilm *generated.Film) *model.FilmModelResponse {
	return &model.FilmModelResponse{
		UUID:        dbFilm.ID,
		Title:       dbFilm.Title,
		Description: dbFilm.Description.String,
		ReleaseDate: dbFilm.Releasedate.Time.Format("2006-01-02"),
		Rate:        int(dbFilm.Rate.Int32),
		Created_at:  dbFilm.CreatedAt,
		Updated_at:  dbFilm.UpdatedAt,
	}
}

func FromDatabaseFilmToFilm(dbFilm *generated.Film) *model.FilmModel {
	return &model.FilmModel{
		UUID:        dbFilm.ID,
		Title:       dbFilm.Title,
		Description: dbFilm.Description.String,
		ReleaseDate: dbFilm.Releasedate.Time.Format("2006-01-02"),
		Rate:        int(dbFilm.Rate.Int32),
		Created_at:  dbFilm.CreatedAt,
		Updated_at:  dbFilm.UpdatedAt,
	}
}

func FromDatabaseFilmsListToFilmsList(dbFilm []generated.Film) []model.FilmsListModel {
	sortedList := make([]model.FilmsListModel, len(dbFilm))
	// прошу прощения за постыдную реализацию, не успел подумать над хорошей
	for i, film := range dbFilm {
		elem := model.FilmsListModel{
			UUID:        film.ID,
			Title:       film.Title,
			Description: film.Description.String,
			ReleaseDate: film.Releasedate.Time.Format("2006-01-02"),
			Rate:        int(film.Rate.Int32),
			Created_at:  film.CreatedAt,
			Updated_at:  film.UpdatedAt,
		}

		sortedList[i] = elem
	}

	return sortedList
}

func FromDatabaseFilmsListByFragmentToFilmsList(dbFilm []generated.Film, fragment string) []model.FilmsListModel {
	sortedList := make([]model.FilmsListModel, len(dbFilm))
	for i, film := range dbFilm {
		if strings.Contains(strings.ToLower(film.Title), strings.ToLower(fragment)) {
			elem := model.FilmsListModel{
				UUID:        film.ID,
				Title:       film.Title,
				Description: film.Description.String,
				ReleaseDate: film.Releasedate.Time.Format("2006-01-02"),
				Rate:        int(film.Rate.Int32),
				Created_at:  film.CreatedAt,
				Updated_at:  film.UpdatedAt,
			}

			sortedList[i] = elem
		}
	}

	return sortedList
}

func FromDatabaseGetActorsListToActorsListByFragment(dbActor []generated.Actor, fragment string) []model.ActorModel {
	sortedList := make([]model.ActorModel, len(dbActor))
	for i, actor := range dbActor {
		if strings.Contains(strings.ToLower(actor.Name), strings.ToLower(fragment)) {
			elem := model.ActorModel{
				UUID:       actor.ID,
				Name:       actor.Name,
				Sex:        actor.Sex,
				Birthday:   actor.Birthday.String(),
				OtherInfo:  actor.Otherinfo.String,
				Created_at: actor.CreatedAt,
				Updated_at: actor.UpdatedAt,
			}

			sortedList[i] = elem
		}
	}

	return sortedList
}

func FromDatabseGetActorsListWithFilmsToService(dblist []generated.GetActorsListWithFilmsRow) *model.GetActorsAndTeirFilmsService {
	// actorsMap := make(map[string][]string)
	//
	// for _, row := range dblist {
	// 	if _, ok := actorsMap[row.ActorName]; !ok {
	// 		actorsMap[row.ActorName] = make([]string, 0)
	// 	}
	// 	actorsMap[row.ActorName] = append(actorsMap[row.ActorName], row.FilmTitle)
	// }
	//
	// service := &model.GetActorsAndTeirFilmsService{
	// 	List: actorsMap,
	// }
	//
	// return service
	service := &model.GetActorsAndTeirFilmsService{
		List: make(map[string][]string),
	}

	for _, row := range dblist {
		if _, ok := service.List[row.ActorName]; !ok {
			service.List[row.ActorName] = make([]string, 0)
		}
		service.List[row.ActorName] = append(service.List[row.ActorName], row.FilmTitle)
	}

	return service
}
