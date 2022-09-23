package helper

import (
	"encoding/json"
	"net/http"
)

func PrintResponseToJson(w http.ResponseWriter, value interface{}) error {
	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(value)
	PrintError(err)

	return err
}
