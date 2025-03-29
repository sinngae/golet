package ginlet

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/sinngae/golet/cmd/servlet/src/handler/ginlet/route"
)

type Service interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type service struct {
	*gin.Engine

	DB *gorm.DB
}

func NewService() Service {
	gin.SetMode(gin.ReleaseMode)
	ng := gin.New()
	// service.Use(middleware)
	route.Init(ng)
	return &service{
		Engine: ng,
	}
}

func (srv *service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	srv.Engine.ServeHTTP(w, req)
}
