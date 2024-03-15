package user

import (
	authjwt "github.com/Alphonnse/filmLibriary/internal/middleware/auth-jwt"
	"github.com/Alphonnse/filmLibriary/internal/repository"
)

type serviceUser struct {
	UserRepository repository.Respository
	AuthService    authjwt.AuthService
}

func NewServiceUser(repo repository.Respository, auth authjwt.AuthService) *serviceUser {
	return &serviceUser{
		UserRepository: repo,
		AuthService:    auth,
	}
}
