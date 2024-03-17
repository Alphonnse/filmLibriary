package filmlibriary

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Alphonnse/filmLibriary/internal/api"
	"github.com/Alphonnse/filmLibriary/internal/convertor"
	"github.com/Alphonnse/filmLibriary/internal/model"
)

// @Summary Add film information
// @Description Add information about a new film
// @ID addFilmInfo
// @Accept json
// @Produce json
// @Param body body model.AddFilmInfoRequest true "Film information to add"
// @Security JWTAdminAuth
// @Success 201 {object} model.FilmModelResponse "Film information added successfully"
// @Failure 400 {object} model.ErrResponse "Invalid JSON or missing fields"
// @Router /film/addFilmInfo [post]
func (i *ImplementationLibriary) AddFilmInfo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		api.RespondWithError(w, 400, fmt.Sprintf("Wrong method: %s", r.Method))
		return
	}

	params := model.AddFilmInfoRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error while decoding add actor info JSON: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if validErrs := ValidateAddFilmInfoRequest(&params); len(validErrs) > 0 {
		log.Printf("Invalid JSON: %v", validErrs)
		api.RespondWithError(w, 400, fmt.Sprintf("Invalid JSON: %v", validErrs))
		return
	}

	createdFilmInfo, err := i.libriaryService.AddFilmInfo(r.Context(), convertor.FromApiAddFilmInfoToService(&params))
	if err != nil {
		log.Printf("Error adding film info: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Couldn't add film info: %v", err))
		return
	}

	api.RespondWithJSON(w, 201, createdFilmInfo)
	log.Printf("New film info is added: %v", params)
}
