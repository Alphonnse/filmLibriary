package user

import "github.com/Alphonnse/filmLibriary/internal/repository"

type serviceUser struct {
	UserRepository repository.Respository
}

func NewServiceUser(repo repository.Respository) *serviceUser {
	return &serviceUser{
		UserRepository: repo,
	}
}
