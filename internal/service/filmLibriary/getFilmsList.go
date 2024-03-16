package filmlibriary

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
)

func (s *serviceLibriary) GetFilmsList(ctx context.Context, info *model.GetFilmsListService) ([]model.FilmsListModel, error) {
	var filmsList []model.FilmsListModel
	var err error

	if info.SortBy == "title" {
		if info.Order == "descending"{
			filmsList, err = s.libriaryRepository.GetFilmsSortedByOrderedByTitleDesc(ctx)
		} else {
			filmsList, err = s.libriaryRepository.GetFilmsSortedByOrderedByTitleAsc(ctx)
		}
	} else if info.SortBy == "releaseDate" {
		if info.Order == "descending"{
			filmsList, err = s.libriaryRepository.GetFilmsSortedByOrderedByReleaseDateDesc(ctx)
		} else {
			filmsList, err = s.libriaryRepository.GetFilmsSortedByOrderedByReleaseDateAsc(ctx)
		}
	} else {
		if info.Order == "ascending"{
			filmsList, err = s.libriaryRepository.GetFilmsSortedByOrderedByRateAsc(ctx)
		} else {
			filmsList, err = s.libriaryRepository.GetFilmsSortedByOrderedByRateDesc(ctx)

		}
	}
	if err != nil {
		return nil, err	
	}

	return filmsList, nil
}
