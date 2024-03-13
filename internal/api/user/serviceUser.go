package user 

import (
	"github.com/Alphonnse/filmLibriary/internal/service"
)


type ImplementationUser struct {
	UserService service.ServiceUserShape
}

func NewImplementationUser(userService service.ServiceUserShape) *ImplementationUser {
	return &ImplementationUser{
		UserService: userService,
	}
}
