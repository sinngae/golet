package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sinngae/gland/cmd/testcase/gin/internal/route/handler"
)

func Init(router *gin.Engine) *gin.Engine {
	SetRoute(router)
	return router
}

func SetRoute(router *gin.Engine) *gin.Engine {
	api := router.Group("/api")
	api.GET("/ping", handler.Ping)
	api.POST("/upload")
	return router
}
