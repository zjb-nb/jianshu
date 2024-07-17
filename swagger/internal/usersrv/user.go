package usersrv

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// get /users/root
func UserHanler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Code": 1000,
		"Data": gin.H{
			"user": gin.H{
				"name": "root",
			},
		},
	})
}
