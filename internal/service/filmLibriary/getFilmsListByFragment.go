package filmlibriary

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
)

func (s *serviceLibriary) GetFilmsListByFragment(ctx context.Context, info *model.GetFilmsListByFragmentService) ([]model.FilmsListModel, error) {
	var filmsList []model.FilmsListModel
	var err error

	if info.FragmentOf == "actor_name" {
		filmsList, err = s.libriaryRepository.GetFilmsByActorNameFragment(ctx, info.Fragment)
	} else {
		filmsList, err = s.libriaryRepository.GetFilmsListByTitleFragment(ctx, info.Fragment)
	}
	if err != nil {
		return nil, err	
	}

	return filmsList, nil
}
