package user

import (
	"app/internal/repo"
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": repo.GetUserLists(),
	})
}
