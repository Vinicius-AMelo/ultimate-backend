package utils

import "net/http"

func WriteReponse(w http.ResponseWriter, body []byte, status int) {
	w.WriteHeader(status)
	w.Write(body)
}
