package controllers_redis

import (
	rabbitmq "api_service/cmd/rabbitmqq"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type session struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
type response struct {
	Message string `json:"message"`
}

func POST(ctx *gin.Context) {
	chann := make(chan string)
	var body session
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	value, err := json.Marshal(body)
	if err != nil {
		ctx.JSON(500, gin.H{"erro": err.Error()})
		return
	}

	rabbitmq.PublishMessage("to_redis", rabbitmq.Message{Key: "setToken", Value: value}, "to_api")

	msg, err := rabbitmq.WaitForResponseMessage("to_api", chann)
	if err != nil {
		ctx.JSON(500, gin.H{"erro": err.Error()})
		return
	}

	var response response
	if err := json.Unmarshal(msg, &response); err != nil {
		ctx.JSON(500, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": msg})

}
