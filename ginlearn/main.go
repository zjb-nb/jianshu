package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	serve := gin.Default()
	serve.GET("home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hi"})
	})
	// serve.Run(":8088")
	http.ListenAndServe(":8088", serve)
}
