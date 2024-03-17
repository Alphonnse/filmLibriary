package filmlibriary

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alphonnse/filmLibriary/internal/api"
)

// @Summary Get list of actors with their associated films
// @Description Retrieve a list of actors along with the films they are associated with
// @ID getActorsInfoListWithFilms
// @Produce json
// @Security JWTRegularUserAuth
// @Success 200 {object} model.GetActorsAndTeirFilmsService "List of actors with associated films"
// @Failure 400 {object} model.ErrResponse "Invalid JSON or missing fields"
// @Router /actor/getList [get]
func (i *ImplementationLibriary) GetActorsInfoListWithFilms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.RespondWithError(w, 400, fmt.Sprintf("Wrong method: %s", r.Method))
		return
	}

	actorsList, err := i.libriaryService.GetActorsListWithFilms(r.Context())
	if err != nil {
		log.Printf("Error returning actors list: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Couldn't return actors list: %v", err))
		return
	}

	api.RespondWithJSON(w, 200, actorsList)
	log.Printf("actors list is returned")
}
