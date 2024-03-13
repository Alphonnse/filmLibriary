package user

import (
	"context"

	"github.com/Alphonnse/filmLibriary/internal/model"
)


func (s *serviceUser) CreateActorInfo(ctx context.Context, info *model.UserModel) error {
	err := s.UserRepository.CreateActorInfo(ctx, info)
	return err 
}
