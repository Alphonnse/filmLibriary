package filmlibriary

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alphonnse/filmLibriary/internal/api"
	"github.com/Alphonnse/filmLibriary/internal/api/validator"
	"github.com/Alphonnse/filmLibriary/internal/convertor"
)

// @Summary Get list of films
// @Description Retrieve a list of films based on specified criteria. It should look like /film/getList/{sortBy}|none|/{orderBy}|none|
// @ID getFilmsList
// @Produce json
// @Security JWTRegularUserAuth
// @Success 200 {object} []model.FilmsListModel "List of films based on criteria"
// @Failure 400 {object} model.ErrResponse "Invalid JSON or missing fields"
// @Router /film/getList/ [get]
func (i *ImplementationLibriary) GetFilmsList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.RespondWithError(w, 400, fmt.Sprintf("Wrong method: %s", r.Method))
		return
	}

	validErrs, sortBy, orderBy := validator.ValidateGetFilmsList(r.URL); 
	if len(validErrs) > 0 {
		log.Printf("Invalid JSON: %v", validErrs)
		api.RespondWithError(w, 406, fmt.Sprintf("Invalid JSON: %v", validErrs))
		return
	}

	sortedFilmsList, err := i.libriaryService.GetFilmsList(r.Context(), convertor.FromApiGetFilmsListToService(sortBy, orderBy))
	if err != nil {
		log.Printf("Error returning sorted films list: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Couldn't return sorted films list: %v", err))
		return
	}

	api.RespondWithJSON(w, 200, sortedFilmsList)
	log.Printf("List is returned")
}
