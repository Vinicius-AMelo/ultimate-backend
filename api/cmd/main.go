package main

import (
	"log"
	"net/http"
	redis_client "ultimate_backend/api/auth"
	redisController "ultimate_backend/api/controllers/redis"
	user_controller "ultimate_backend/api/controllers/user"
	UserModel "ultimate_backend/api/models/user"

	"ultimate_backend/api/database"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	log.Println("Trying to connect to redis")
	redis_client.InitClient()
	log.Println("Connected to redis")

	log.Println("Trying to connect to postgres")
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to postgres")
	createTables()
	defer db.Close()

	r.HandleFunc("/session", redisController.POST).Methods(http.MethodPost)
	r.HandleFunc("/session", redisController.GET).Methods(http.MethodGet)

	r.HandleFunc("/user", user_controller.POST).Methods(http.MethodPost)
	r.HandleFunc("/user", user_controller.GetAll).Methods(http.MethodGet)

	log.Println("Server running in port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createTables() {
	UserModel.CreateUserTable()
}
