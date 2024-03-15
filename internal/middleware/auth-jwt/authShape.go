package authjwt

import (
	"context"
	"net/http"

	"github.com/Alphonnse/filmLibriary/internal/model"
)

type AuthService interface {
	// MiddlewareAuth(func(http.ResponseWriter, *http.Request)) http.HandlerFunc
	JWTAdminAuth(func(http.ResponseWriter, *http.Request)) http.HandlerFunc
	JWTRegularUserAuth(func(http.ResponseWriter, *http.Request)) http.HandlerFunc
	GenerateJWT(ctx context.Context, user *model.UserModel) (string, error)
	// GenerateJWT(context.Context, model.UserModel) (string, error)
}
