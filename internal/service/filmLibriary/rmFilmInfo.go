package filmlibriary

import (
	"context"
	"errors"

	"github.com/google/uuid"
)


func (s *serviceLibriary) RmFilmInfo(ctx context.Context, uuid uuid.UUID ) (error) {
	_, err := s.libriaryRepository.GetFilmByID(ctx, uuid)
	if err == nil {
		return s.libriaryRepository.RmFilmInfo(ctx, uuid)
	}
	return errors.New("No film with what uuid in DB")
}
