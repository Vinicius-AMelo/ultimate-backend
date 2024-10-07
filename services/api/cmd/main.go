package main

import (
	controllers_redis "api_service/cmd/controllers"
	rabbitmq "api_service/cmd/rabbitmqq"

	"github.com/gin-gonic/gin"
)

func main() {
	rabbitmq.InitQueue("to_redis")
	rabbitmq.InitQueue("to_api")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/session", controllers_redis.POST)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
