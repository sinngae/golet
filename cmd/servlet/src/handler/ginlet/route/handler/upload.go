package handler

import "github.com/gin-gonic/gin"

func Upload(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
