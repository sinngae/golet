package route

import (
	"github.com/gin-gonic/gin"
<<<<<<<< HEAD:cmd/servlet/src/handler/ginlet/route/route.go
	"github.com/sinngae/golet/cmd/servlet/src/handler/ginlet/route/handler"
========
	"github.com/sinngae/golet/sample/server/internal/gin/route/handler"
>>>>>>>> 4aad3e1b64427d5ebafb07f037b140f7eb3a6511:sample/server/internal/gin/route/route.go
)

func Init(router *gin.Engine) *gin.Engine {
	SetRoute(router)
	return router
}

func SetRoute(router *gin.Engine) *gin.Engine {
	api := router.Group("/api")
	api.GET("/ping", handler.Ping)
<<<<<<<< HEAD:cmd/servlet/src/handler/ginlet/route/route.go
========
	api.POST("/upload")

	_ = api.Group("/proxy", handler.Proxy)
>>>>>>>> 4aad3e1b64427d5ebafb07f037b140f7eb3a6511:sample/server/internal/gin/route/route.go
	return router
}
