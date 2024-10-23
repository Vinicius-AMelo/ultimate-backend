package main

import (
	"log"
	"net/http"
	redis_client "ultimate_backend/api/auth"
	redisController "ultimate_backend/api/controllers/redis"
	userController "ultimate_backend/api/controllers/user"
	UserModel "ultimate_backend/api/models/user"

	"ultimate_backend/api/database"
)

func main() {

	r := http.NewServeMux()
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

	redisHandler := &redisController.RedisHandler{}
	userHandler := &userController.UserHandler{}

	r.HandleFunc("POST /session", redisHandler.Post)
	r.HandleFunc("GET /session", redisHandler.Get)

	r.HandleFunc("POST /user", userHandler.Post)
	r.HandleFunc("GET /user", userHandler.GetAll)
	r.HandleFunc("POST /auth", userHandler.HandleLogin)

	log.Println("Server running in port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createTables() {
	UserModel.CreateUserTable()
}
