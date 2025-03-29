package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sinngae/golet/cmd/servlet/src/handler/ginlet/route/handler"
)

func Init(router *gin.Engine) *gin.Engine {
	SetRoute(router)
	return router
}

func SetRoute(router *gin.Engine) *gin.Engine {
	api := router.Group("/api")
	api.GET("/ping", handler.Ping)
	api.POST("/upload")

	_ = api.Group("/proxy", handler.Proxy)
	return router
}
