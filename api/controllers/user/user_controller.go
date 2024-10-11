package user_controller

import (
	"encoding/json"
	"io"
	"net/http"
	UserModel "ultimate_backend/api/models/user"
	"ultimate_backend/api/utils"
)

func POST(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteReponse(w, []byte(err.Error()), http.StatusInternalServerError)
		return
	}

	var user UserModel.User
	if err := json.Unmarshal(body, &user); err != nil {
		utils.WriteReponse(w, []byte(err.Error()), http.StatusInternalServerError)
		return
	}

	if err := UserModel.InsertUser(user); err != nil {
		utils.WriteReponse(w, []byte(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.WriteReponse(w, []byte("User created"), http.StatusOK)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := UserModel.GetUsers()
	if err != nil {
		utils.WriteReponse(w, []byte(err.Error()), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		utils.WriteReponse(w, []byte(err.Error()), http.StatusInternalServerError)
		return
	}
}
