package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonNewEncoder(w http.ResponseWriter, statusCode int, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Fatal(err)
	}
}
