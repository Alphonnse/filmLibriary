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

func (i *ImplementationLibriary) AddActorInfo(w http.ResponseWriter, r *http.Request) {

	// if r.Method != http.MethodPost {
	// 	api.RespondWithError(w, 400, fmt.Sprintf("Wrong method: %s", r.Method))
	// 	return
	// }

	params := model.AddActorInfoRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error while decoding add actor info JSON: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if validErrs := ValidateActorAddRequest(&params); len(validErrs) > 0 {
		log.Printf("Invalid JSON: %v", validErrs)
		api.RespondWithError(w, 400, fmt.Sprintf("Invalid JSON: %v", validErrs))
		return
	}

	createdActorInfo, err := i.libriaryService.AddActorInfo(r.Context(), convertor.FromApiAddActorInfoToService(&params))
	if err != nil {
		log.Printf("Error adding actor info: %v\n", err)
		api.RespondWithError(w, 400, fmt.Sprintf("Couldn't add actor info: %v", err))
		return
	}

	api.RespondWithJSON(w, 201, createdActorInfo)
	log.Printf("New actor is added: %s", params)
}
