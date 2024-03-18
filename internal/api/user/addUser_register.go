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

// registration of user and adding user
// @Summary Register a new user
// @Description Register a new user with the provided details
// @ID registerUser
// @Accept json
// @Produce json
// @Param body body model.Register true "User registration details"
// @Success 201 {object} model.UserRequestModel "User registered successfully"
// @Failure 400 {object} model.ErrResponse "Invalid JSON or missing fields"
// @Failure 500 {object} model.ErrResponse "Failed to register user"
// @Router /register [post]
func (i *ImplementationUser) AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		api.RespondWithError(w, 400, fmt.Sprintf("Wrong method: %s", r.Method))
		return
	}

	params := model.Register{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	
	if err != nil {
		log.Printf("Error while decoding register JSON: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if validErrs := validator.ValidateUserAddRequest(&params); len(validErrs) > 0 {
		log.Printf("Invalid JSON: %v", validErrs)
		api.RespondWithError(w, 400, fmt.Sprintf("Invalid JSON: %v", validErrs))
		return
	}

	addUserInfo, err := i.UserService.AddUser(r.Context(), convertor.FromApiAddUserService(&params))
	if err != nil {
		log.Printf("Error registering user: %v\n", err)
		api.RespondWithError(w, 401, fmt.Sprintf("Couldn't registering user : %v", err))
		return
	}

	api.RespondWithJSON(w, 201, addUserInfo)
	log.Printf("user %s is registered", params.Name)
}
func (i *ImplementationUser) Test(w http.ResponseWriter, r *http.Request) {
	api.RespondWithJSON(w, 201, "all rigth")
}
