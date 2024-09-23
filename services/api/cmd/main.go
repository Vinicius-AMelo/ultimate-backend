package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	rabbitmq "github.com/my/repo/services/rabbitmq"
)

type session struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func main() {
	rabbitmq.InitQueue("redis")
	// rabbitmq.InitQueue("responseToApi")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/session", func(ctx *gin.Context) {
		var body session
		err := ctx.BindJSON(&body)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		value, err := json.Marshal(body)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		rabbitmq.PublishMessage("redis", rabbitmq.Message{Key: "setToken", Value: value})

		ctx.JSON(200, gin.H{"id": body.ID, "token": body.Token})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
