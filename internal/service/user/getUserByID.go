package user

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/google/uuid"
)

func (s *serviceUser) GetUserByID(ctx context.Context, uuid uuid.UUID) (*model.UserModel, error) {
	userInfo, err := s.UserRepository.GetUserByID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
