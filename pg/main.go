package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	engine := gin.Default()

	engine.GET("/hello", authMiddleware(), handlerHello())

	personas := engine.Group("/personas", authMiddleware())

	{
		personas.GET("", getAllPersonas)
		personas.POST("", savePersonas)
	}

	engine.Run(":8080")
}

func getAllPersonas(ctx *gin.Context) {
	ctx.JSON(200, "GET PERSONAS")
}
func savePersonas(ctx *gin.Context) {
	ctx.JSON(200, "POST PERSONAS")
}

func handlerHello() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message":"hello world"})
	}
}

func authMiddleware() gin.HandlerFunc {
	token := "1234"

	return func(ctx *gin.Context) {
		tokenReq := ctx.GetHeader("token")
		if tokenReq != token {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error":"Invalid token"})
			return
		}

		ctx.Writer.Written();

		ctx.Next()

		fmt.Println("Despues")
	}
}


