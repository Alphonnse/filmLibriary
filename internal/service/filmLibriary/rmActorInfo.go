package filmlibriary

import (
	"context"
	"errors"

	"github.com/google/uuid"
)


func (s *serviceLibriary) RmActorInfo(ctx context.Context, uuid uuid.UUID) (error) {
	_, err := s.libriaryRepository.GetActorByID(ctx, uuid)
	if err == nil {
		return s.libriaryRepository.RmActorInfo(ctx, uuid)
	}
	return errors.New("No actor with that uuid in DB")
}
