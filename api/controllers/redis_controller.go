package redisController

import (
	redisClient "ultimate_backend/api/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type session struct {
	ID string `json:"id"`
}

func POST(ctx *gin.Context) {
	var body session

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token := uuid.New().String()

	if err := redisClient.SetToken(ctx.Request.Context(), body.ID, token); err != nil {
		ctx.JSON(500, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Token created", "token": token})

}

func GetToken(ctx *gin.Context) {
	var body session

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	token := redisClient.GetToken(ctx.Request.Context(), body.ID)

	if token == "" {
		ctx.JSON(404, gin.H{"error": "Token not found"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Token found", "token": token})

}
