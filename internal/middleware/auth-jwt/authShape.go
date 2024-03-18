package authjwt

import (
	"context"
	"net/http"

	"github.com/Alphonnse/filmLibriary/internal/model"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name AuthService 
type AuthService interface {
	JWTAdminAuth(func(http.ResponseWriter, *http.Request)) http.HandlerFunc
	JWTRegularUserAuth(func(http.ResponseWriter, *http.Request)) http.HandlerFunc
	GenerateJWT(ctx context.Context, user *model.UserModel) (string, error)
}
