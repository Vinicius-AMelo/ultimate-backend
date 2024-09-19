package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

type session struct {
	ID string `json:"id"`
	Token string `json:"token"`
}

func main(){
	r := gin.Default()

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/");
	if err != nil{
		panic(err)
	} 
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	r.POST("/session", func(ctx *gin.Context) {
		var body session
		err := ctx.BindJSON(&body); if err != nil {
			return
		}
	})
}