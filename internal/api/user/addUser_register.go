package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Alphonnse/filmLibriary/internal/api"
	"github.com/Alphonnse/filmLibriary/internal/convertor"
	"github.com/Alphonnse/filmLibriary/internal/model"
)

// registration of user and adding user 
func (i *ImplementationUser) AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		api.RespondWithError(w, 400, fmt.Sprintf("Wrong method: %s", r.Method))
		return
	}

	params := model.Register{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	
	if err != nil {
		log.Printf("Error while decoding add user JSON: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if validErrs := ValidateUserAddRequest(&params); len(validErrs) > 0 {
		log.Printf("Invalid JSON: %v", validErrs)
		api.RespondWithError(w, 400, fmt.Sprintf("Invalid JSON: %v", validErrs))
		return
	}

	addUserInfo, err := i.UserService.AddUser(r.Context(), convertor.FromApiAddUserService(&params))
	if err != nil {
		log.Printf("Error adding user: %v\n", err)
		api.RespondWithError(w, 500, fmt.Sprintf("Couldn't add user : %v", err))
		return
	}

	api.RespondWithJSON(w, 201, addUserInfo)
	log.Printf("user %s is registered", params.Name)
}
