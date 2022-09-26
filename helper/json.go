package helper

import (
	"encoding/json"
	"net/http"
)

//func PrintResponseToJson(w http.ResponseWriter, value interface{}) error {
//	w.Header().Set("content-type", "application/json")
//	encoder := json.NewEncoder(w)
//	err := encoder.Encode(value)
//	PrintError(err)
//
//	return err
//}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	encoder := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := encoder.Encode(payload)
	PrintError(err)
}
