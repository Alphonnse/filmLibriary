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
