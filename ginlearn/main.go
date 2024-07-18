package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	serve := gin.Default()

	serve.Use(func(ctx *gin.Context) {
		fmt.Println(1)
		ctx.Next()
		fmt.Println(4)
	}, func(ctx *gin.Context) {
		fmt.Println(2)
		// ctx.Abort()
		ctx.Next()
		fmt.Println(3)
	})

	serve.GET("home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hi"})
	})

	v1 := serve.Group("/v1")
	{
		v1.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "hi"})
		})
	}
	// serve.Run(":8088")
	http.ListenAndServe(":8088", serve)
}
