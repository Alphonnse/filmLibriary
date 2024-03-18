package userauth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Alphonnse/filmLibriary/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (a *auth) GenerateJWT(ctx context.Context, user *model.UserModel) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.UUID,
		"role": user.Role_id,
		"iat":  time.Now().Unix(),
		"eat":  time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(a.userAuth.Secret())
}

func (a *auth) ValidateJWT(w http.ResponseWriter, r *http.Request) error {
	token, err := a.getToken(w, r)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token") 
}

func (a *auth) ValidateAdminRoleJwt(w http.ResponseWriter, r *http.Request) error {
	token, err := a.getToken(w, r)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 1{
		return nil
	}
	return errors.New("invalid admin token provided token") 
}

func (a *auth) ValidateRegularUesrRoleJWT(w http.ResponseWriter, r *http.Request) error {
	token, err := a.getToken(w, r)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 2 || userRole == 1{
		return nil
	}
	return errors.New("invalid admin token provided token") 
}

func (a *auth) CurrentUser(ctx context.Context, w http.ResponseWriter, r *http.Request) *model.UserModel {
	err := a.ValidateJWT(w, r)
	if err != nil {
		return nil 
	}
	token, _ := a.getToken(w, r)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId, _ := uuid.Parse(claims["id"].(string))

	user, err := a.UserService.GetUserByID(ctx, userId)
	if err != nil {
		return nil
	}

	return user
}

func (a *auth) getToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenStringFromRequest, err := r.Cookie("Authorization")
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenStringFromRequest.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.userAuth.Secret(), nil

	})

	return token, err
}
