package filmlibriary

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Alphonnse/filmLibriary/internal/api"
	"github.com/Alphonnse/filmLibriary/internal/convertor"
	"github.com/google/uuid"
)

func (i *ImplementationLibriary) RmActorInfo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		api.RespondWithError(w, 400, fmt.Sprintf("Wrong method: %s", r.Method))
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	id := urlParts[len(urlParts)-1]
	actorID, err := uuid.Parse(id)

	if err != nil {
		log.Printf("Invalid actor UUID: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Invalid actor UUID: %v\n", err))
		return
	}

	err = i.libriaryService.RmActorInfo(r.Context(), convertor.FromApiRmActorInfoToService(actorID))
	if err != nil {
		log.Printf("Error removing actor info: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Couldn't remove actor info: %v", err))
		return
	}

	api.RespondWithJSON(w, 201, fmt.Sprintf("Actor with UUID %s removed", actorID))
	log.Printf("Actor with UUID %s removed", actorID)
}
