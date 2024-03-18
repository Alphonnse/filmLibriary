package userauth_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Alphonnse/filmLibriary/internal/config"
	userauth "github.com/Alphonnse/filmLibriary/internal/middleware/auth-jwt/userAuth"
	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/Alphonnse/filmLibriary/internal/service/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	mockService := new(mocks.ServiceUserShape)

	err := config.Load("./../../../../.env")
	if err != nil {
		t.Logf("shoead: %s", err)
	}
	jwtConf, err := config.NewJwtConfig()
	assert.NoError(t, err)

	newAuth := userauth.NewAuthService(mockService, jwtConf)

	// first testcase with wrong token
	wrongReqCookie := httptest.NewRequest(http.MethodPost, "/actor/remove/asdf", bytes.NewBuffer([]byte{}))
	w := httptest.NewRecorder()

	wrongReqCookie.AddCookie(createCookieUsingToken("asd"))

	err = newAuth.ValidateJWT(w, wrongReqCookie)
	assert.Error(t, err)

	// user to create a token
	mockUserModel := model.UserModel{
		UUID:       uuid.New(),
		Role_id:    2,
		Name:       "Vanya",
		Created_at: time.Now().UTC(),
		Updated_at: time.Now().UTC(),
		Password:   "qwer",
	}

	// generating the token for user
	token, err := newAuth.GenerateJWT(context.Background(), &mockUserModel)
	assert.NoError(t, err)

	RighReqCookie := httptest.NewRequest(http.MethodPost, "/actor/remove/asdf", bytes.NewBuffer([]byte{}))
	RighReqCookie.AddCookie(createCookieUsingToken(token))

	// Test JWT validation
	err = newAuth.ValidateJWT(w, RighReqCookie)
	assert.NoError(t, err)

	// Test validation of admin role with no access
	err = newAuth.ValidateAdminRoleJwt(w, RighReqCookie)
	assert.Error(t, err)

	// Test validation of regular user role with access 
	err = newAuth.ValidateRegularUesrRoleJWT(w, RighReqCookie)
	assert.NoError(t, err)


	mockService.AssertExpectations(t)
}

func createCookieUsingToken(token string) *http.Cookie{
	return &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		MaxAge:   3600 * 24 * 30,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteDefaultMode,
	}
}
