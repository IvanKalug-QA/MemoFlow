package res

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, data any, statusCode int) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
}
