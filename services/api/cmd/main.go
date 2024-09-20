package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	rabbitmq "github.com/my/repo/services/api/messages"
)

type session struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}



func main() {
	rabbitmq.InitQueue()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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

		rabbitmq.PublishMessage(context, rabbitmq.Message{Key: "redis", Value: value})

		ctx.JSON(200, gin.H{"id": body.ID, "token": body.Token})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
