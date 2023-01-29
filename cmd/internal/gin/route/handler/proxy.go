package handler

import "github.com/gin-gonic/gin"

func Proxy(ctx *gin.Context) {
	// rewrite
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
