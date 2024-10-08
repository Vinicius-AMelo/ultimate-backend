package main

import (
	"log"
	redis_client "ultimate_backend/api/auth"
	redis_controller "ultimate_backend/api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	log.Println("Trying to connect to redis")
	redis_client.InitClient()
	log.Println("Connected to redis")

	r.POST("/session", redis_controller.POST)
	r.GET("/session", redis_controller.GetToken)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
