package filmlibriary

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
)


func (s *serviceLibriary) RmFilmInfo(ctx context.Context, model *model.FilmModel) (error) {
	return s.libriaryRepository.RmFilmInfo(ctx, model.UUID)
}
