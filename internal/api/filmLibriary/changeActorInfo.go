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

func (i *ImplementationLibriary) ChangeActorInfo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		api.RespondWithError(w, 400, fmt.Sprintf("Wrong method: %s", r.Method))
		return
	}

	params := model.ChangeActorInfoRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error while decoding change actor info JSON: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if validErrs := ValidateChangeActorRequest(&params); len(validErrs) > 0 {
		log.Printf("Invalid JSON: %v", validErrs)
		api.RespondWithError(w, 400, fmt.Sprintf("Invalid JSON: %v", validErrs))
		return
	}

	changedActorInfo, err := i.libriaryService.ChangeActorInfo(r.Context(), convertor.FromApiChangeActorInfoToService(&params))
	if err != nil {
		log.Printf("Error change actor info: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Couldn't change actor info: %v", err))
		return
	}

	api.RespondWithJSON(w, 201, changedActorInfo)
	log.Printf("Info of actor %s is changed to: %s", params.UUID, params)
}
