package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alphonnse/filmLibriary/internal/model"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}
	RespondWithJSON(w, code, model.ErrResponse{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
