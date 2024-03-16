package filmlibriary

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
)

func (s *serviceLibriary) GetActorsListWithFilms(ctx context.Context) (*model.GetActorsAndTeirFilmsService, error) {
	actors, err := s.libriaryRepository.GetActorsListWithFilms(ctx)
	if err != nil {
		return nil, err	
	}

	return actors, nil
}
