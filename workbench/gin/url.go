package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(ng *gin.Engine) *gin.Engine {
	root := ng.Group("/api")

	test := root.Group("test", func(ctx *gin.Context) {

		ctx.Next()
	})
	test.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	return ng
}
