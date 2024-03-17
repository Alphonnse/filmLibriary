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

// @Summary Get list of films by fragment
// @Description Retrieve a list of films based on a search fragment
// @ID getFilmsListByFragment
// @Accept json
// @Produce json
// @Param body body model.GetFilmsListByFragmentRequest true "Search fragment for films"
// @Security JWTRegularUserAuth
// @Success 200 {object} []model.FilmsListModel "List of films matching the search fragment"
// @Failure 400 {object} model.ErrResponse "Invalid JSON or missing fields"
// @Router /film/findFilm [get]
func (i *ImplementationLibriary) GetFilmsListByFragment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.RespondWithError(w, 400, fmt.Sprintf("Wrong method: %s", r.Method))
		return
	}

	params := model.GetFilmsListByFragmentRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error while decoding get films list request JSON: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if validErrs := ValidateGetFilmsListByFragment(&params); len(validErrs) > 0 {
		log.Printf("Invalid JSON: %v", validErrs)
		api.RespondWithError(w, 400, fmt.Sprintf("Invalid JSON: %v", validErrs))
		return
	}

	sortedFilmsList, err := i.libriaryService.GetFilmsListByFragment(r.Context(), convertor.FromApiGetFilmsListByFragmentToService(&params))
	if err != nil {
		log.Printf("Error returning sorted films list: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Couldn't return sorted films list: %v", err))
		return
	}

	api.RespondWithJSON(w, 200, sortedFilmsList)
	log.Printf("List is returned: %s", params)
}
