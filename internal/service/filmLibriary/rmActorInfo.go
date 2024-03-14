package filmlibriary

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
)


func (s *serviceLibriary) RmActorInfo(ctx context.Context, model *model.ActorModel) (error) {
	return s.libriaryRepository.RmActorInfo(ctx, model.UUID)
}
