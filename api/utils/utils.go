package utils

import (
	"encoding/json"
	"net/http"
)

func WriteReponse(w http.ResponseWriter, body []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(body)
}

func WriteReponseJson(w http.ResponseWriter, body any) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		return err
	}
	return nil
}
