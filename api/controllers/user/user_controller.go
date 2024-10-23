package userController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	UserModel "ultimate_backend/api/models/user"
	"ultimate_backend/api/utils"
)

type UserHandler struct{}

func (h *UserHandler) Post(w http.ResponseWriter, r *http.Request) {
	log.Println("request received")
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

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Println("request received")
	users, err := UserModel.GetUsers()
	if err != nil {
		utils.WriteReponse(w, []byte(err.Error()), http.StatusInternalServerError)
		return
	}

	if err := utils.WriteReponseJson(w, users); err != nil {
		utils.WriteReponse(w, []byte(err.Error()), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
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
	storedUser, err := UserModel.GetUser(user.Email)
	if err != nil {
		utils.WriteReponse(w, []byte(err.Error()), http.StatusInternalServerError)
		return
	}
	log.Println(storedUser.Password)
	log.Println(user.Password)
	if storedUser.Password == user.Password {
		utils.WriteReponse(w, []byte(`{"message": "OK"`), http.StatusOK)
		log.Println(200)

	} else {
		utils.WriteReponse(w, []byte(`{"message": "Unauthorized"}`), http.StatusUnauthorized)
		log.Println(401)
	}

}
