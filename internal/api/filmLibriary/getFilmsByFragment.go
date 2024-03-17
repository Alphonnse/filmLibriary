package filmlibriary

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alphonnse/filmLibriary/internal/api"
	"github.com/Alphonnse/filmLibriary/internal/convertor"
)

// @Summary Get list of films by fragment
// @Description Retrieve a list of films based on a search fragment. It should look like /film/findFilm/{fragmentOf}/{fragment}
// @ID getFilmsListByFragment
// @Produce json
// @Security JWTRegularUserAuth
// @Success 200 {object} []model.FilmsListModel "List of films matching the search fragment"
// @Failure 400 {object} model.ErrResponse "Invalid JSON or missing fields"
// @Router /film/findFilm [get]
func (i *ImplementationLibriary) GetFilmsListByFragment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.RespondWithError(w, 400, fmt.Sprintf("Wrong method: %s", r.Method))
		return
	}

	validErrs, fragmentOf, fragment := ValidateGetFilmsListByFragment(r.URL)
	if len(validErrs) > 0 {
		log.Printf("Invalid JSON: %v", validErrs)
		api.RespondWithError(w, 406, fmt.Sprintf("Invalid JSON: %v", validErrs))
		return
	}

	sortedFilmsList, err := i.libriaryService.GetFilmsListByFragment(r.Context(), convertor.FromApiGetFilmsListByFragmentToService(fragmentOf, fragment))
	if err != nil {
		log.Printf("Error returning sorted films list: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Couldn't return sorted films list: %v", err))
		return
	}

	api.RespondWithJSON(w, 200, sortedFilmsList)
	log.Printf("List is returned")
}
