package helpers

import (
	"encoding/json"
	"net/http"
)

func FormatResponse (w http.ResponseWriter ,  data ServiceResponse) {
	w.WriteHeader(data.StatusCode)
	json.NewEncoder(w).Encode(data)
}
