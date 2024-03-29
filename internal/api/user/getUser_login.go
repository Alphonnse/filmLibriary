package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Alphonnse/filmLibriary/internal/api"
	"github.com/Alphonnse/filmLibriary/internal/api/validator"
	"github.com/Alphonnse/filmLibriary/internal/convertor"
	"github.com/Alphonnse/filmLibriary/internal/model"
)

//login

// @Summary Login as a user
// @Description Log in with user credentials to get access token
// @ID loginUser
// @Accept json
// @Produce json
// @Param body body model.Login true "User login credentials"
// @Success 201 {object} model.UserRequestModel "User logged in successfully"
// @Failure 400 {object} model.ErrResponse "Invalid JSON or missing fields"
// @Failure 500 {object} model.ErrResponse "Failed to log in user"
// @Router /login [post]
func (i *ImplementationUser) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		api.RespondWithError(w, 400, fmt.Sprintf("Wrong method: %s", r.Method))
		return
	}

	params := model.Login{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	
	if err != nil {
		log.Printf("Error while decoding login JSON: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if validErrs := validator.ValidateUserGetRequest(&params); len(validErrs) > 0 {
		log.Printf("Invalid JSON: %v", validErrs)
		api.RespondWithError(w, 400, fmt.Sprintf("Invalid JSON: %v", validErrs))
		return
	}

	getUserInfo , tokenString, err := i.UserService.GetUser(r.Context(), convertor.FromApiGetUserService(&params))
	if err != nil {
		log.Printf("Error getting user: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Couldn't get user : %v", err))
		return
	}

	cookie := http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		MaxAge:   3600 * 24 * 30,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(w, &cookie)

	api.RespondWithJSON(w, 201, getUserInfo)
	log.Printf("user %s is loggined in", params.Name)
}
