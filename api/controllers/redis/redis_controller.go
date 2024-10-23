package redisController

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	redisClient "ultimate_backend/api/auth"

	"github.com/google/uuid"
)

type session struct {
	ID string `json:"id"`
}

type response struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type RedisHandler struct {
}

func (h *RedisHandler) Post(w http.ResponseWriter, r *http.Request) {
	var session session

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := json.Unmarshal(body, &session); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	token := uuid.New().String()

	if err := redisClient.SetToken(r.Context(), session.ID, token); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(response{Message: "Token created", Token: token}); erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	}
}

func (h *RedisHandler) Get(w http.ResponseWriter, r *http.Request) {
	var session session

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := json.Unmarshal(body, &session); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	token := redisClient.GetToken(r.Context(), session.ID)

	if token == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Token not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"message": "Token found", "token": "%s"}`, token)))
}
