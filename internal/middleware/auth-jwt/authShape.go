package authjwt

import "net/http"

type AuthService interface {
	MiddlewareAuth(func(http.ResponseWriter, *http.Request)) http.HandlerFunc
}
