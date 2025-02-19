package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sinngae/golet/sample/server/internal/gin/route/handler/internal/proxy"
)

func Proxy(ctx *gin.Context) {
	// rewrite
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
