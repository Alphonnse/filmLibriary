package userauth

import (
	"fmt"
	"net/http"

	"github.com/Alphonnse/filmLibriary/internal/api"
)

func (a *auth) JWTAdminAuth(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := a.ValidateJWT(w, r)
		if err != nil {
			api.RespondWithError(w, 401, fmt.Sprintf("Authorization required: %v", err))
			return
		}
		err = a.ValidateAdminRoleJwt(w, r)
		if err != nil {
			api.RespondWithError(w, 401, fmt.Sprintf("Only Administrator is allowed to perform this action: %v", err))
			return
		}
		handler(w, r)
	}
}

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
			return
		}
		handler(w, r)
	}
}

// func (a *auth) MiddlewareAuth(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		tokenString, err := r.Cookie("Authorization")
// 		if err != nil {
// 			api.RespondWithError(w, 401, fmt.Sprint("Error with cookie: ", err))
// 			return
// 		}
//
// 		token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return a.userAuth.Secret(), nil
//
// 		})
//
// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//
// 			if float64(time.Now().Unix()) > claims["exp"].(float64) {
// 				api.RespondWithError(w, 401, fmt.Sprint("Error with JWT: ", err))
// 			}
//
// 			subUUID, err := uuid.Parse(claims["sub"].(string))
// 			if err != nil {
// 				api.RespondWithError(w, 500, fmt.Sprint("Error while parsing UUID: ", err))
// 				return
// 			}
//
// 			_, err = a.UserService.GetUserByID(r.Context(), subUUID)
// 			if err != nil {
// 				api.RespondWithError(w, 500, fmt.Sprint("No matched UUID in DB: ", err))
// 			}
//
// 			// Attach to req and continue
// 			handler(w, r)
//
// 		} else {
// 			api.RespondWithError(w, 401, fmt.Sprint("Error with JWT: ", err))
// 		}
// 	}
// }
