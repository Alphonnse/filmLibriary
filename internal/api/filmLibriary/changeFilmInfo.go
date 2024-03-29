package filmlibriary

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

// @Summary Change film information
// @Description Update information about an existing film
// @ID changeFilmInfo
// @Accept json
// @Produce json
// @Param body body model.ChangeFilmInfoRequest true "Film information to change"
// @Security JWTAdminAuth
// @Success 201 {object} model.FilmModelResponse "Film information changed successfully"
// @Failure 400 {object} model.ErrResponse "Invalid JSON or missing fields"
// @Router /film/changeFilmInfo [post]
func (i *ImplementationLibriary) ChangeFilmInfo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		api.RespondWithError(w, 400, fmt.Sprintf("Wrong method: %s", r.Method))
		return
	}

	params := model.ChangeFilmInfoRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error while decoding change actor info JSON: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if validErrs := validator.ValidateChangeFilmInfoRequest(&params); len(validErrs) > 0 {
		log.Printf("Invalid JSON: %v", validErrs)
		api.RespondWithError(w, 400, fmt.Sprintf("Invalid JSON: %v", validErrs))
		return
	}

	changedFilmInfo, err := i.libriaryService.ChangeFilmInfo(r.Context(), convertor.FromApiChangeFilmInfoToService(&params))
	if err != nil {
		log.Printf("Error change film info: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Couldn't change film info: %v", err))
		return
	}

	api.RespondWithJSON(w, 201, changedFilmInfo)
	log.Printf("The information about the film %v has been changed", params.UUID)
}
