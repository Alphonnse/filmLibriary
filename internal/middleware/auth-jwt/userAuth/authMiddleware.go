package userauth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alphonnse/filmLibriary/internal/api"
)

// @securityDefinitions.apikey JWTAdminAuth 
// @in header
// @name AuthorizationAdmin
func (a *auth) JWTAdminAuth(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := a.ValidateJWT(w, r)
		if err != nil {
			api.RespondWithError(w, 401, fmt.Sprintf("Authorization required: %v", err))
			return
		}
		err = a.ValidateAdminRoleJwt(w, r)
		if err != nil {
			api.RespondWithJSON(w, 401, fmt.Sprintf("Error authorization: %v", err))
			return
		}
		handler(w, r)
	}
}

// @securityDefinitions.apikey JWTRegularUserAuth
// @in header
// @name AuthorizationRegularUser
func (a *auth) JWTRegularUserAuth(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := a.ValidateJWT(w, r)
		if err != nil {
			api.RespondWithError(w, 401, fmt.Sprintf("Authorization required: %v", err))
			return
		}
		err = a.ValidateRegularUesrRoleJWT(w, r)
		if err != nil {
			api.RespondWithError(w, 401, fmt.Sprintf("Only Administrator is allowed to perform this action: %v", err))
			log.Println("Admission without permission")
			return
		}
		handler(w, r)
	}
}

