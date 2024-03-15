package userauth

import (
	"github.com/Alphonnse/filmLibriary/internal/config"
	"github.com/Alphonnse/filmLibriary/internal/service"
)

type auth struct {
	UserService service.ServiceUserShape
	userAuth    config.JwtConfig
}

func NewAuthService(service service.ServiceUserShape, config config.JwtConfig) *auth {
	secret := config
	return &auth{
		UserService: service,
		userAuth:    secret,
	}
}
