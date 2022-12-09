package gin

import (
	"context"
	"fmt"
	"github.com/sinngae/golet/cmd/golet/internal/gin/midware"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sinngae/golet/cmd/golet/internal/gin/route"
)

type Server struct {
	*gin.Engine
	entity *http.Server

	port string
}

func NewServer(port string) *Server {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	// common middleware
	pprof.Register(engine)
	engine.Use(midware.Recovery())
	// http req route
	route.Init(engine)

	return &Server{
		Engine: engine,
		port:   port,
	}
}

func (srv *Server) Startup() error {
	server := &http.Server{
		Addr:    srv.port,
		Handler: srv,
	}
	srv.entity = server
	fmt.Printf("server listening at %s ...\n", srv.port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		err = fmt.Errorf("server serve failed, err=%v\n", err)
		return err
	}
	// quit
	fmt.Printf("server quiting @%s\n", time.Now())
	return nil
}

func (srv *Server) Shutdown(ctx context.Context) error {
	if err := srv.entity.Shutdown(ctx); err != nil {
		fmt.Printf("server shutdown failed, err=%s", err)
		return err
	}
	return nil
}
