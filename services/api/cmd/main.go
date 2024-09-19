package main

import (
	"github.com/gin-gonic/gin"
)

type session struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/");
	// if err != nil{
	// 	panic(err)
	// }
	// defer conn.Close()

	// ch, err := conn.Channel()
	// if err != nil {
	// 	panic(err)
	// }
	// defer ch.Close()

	r.POST("/session", func(ctx *gin.Context) {
		var body session
		err := ctx.BindJSON(&body)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"id": body.ID, "token": body.Token})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
