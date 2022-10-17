package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(ng *gin.Engine) *gin.Engine {
	root := ng.Group("/api")

	test := root.Group("test", func(ctx *gin.Context) {
		ctx.Set("a", 2)
		ctx.Next()
	}, gin.Recovery())
	test.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	return ng
}
