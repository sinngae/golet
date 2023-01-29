package gin

import (
	"context"
	"fmt"
	"github.com/sinngae/golet/cmd/internal/gin/midware"
	"github.com/sinngae/golet/cmd/internal/gin/route"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
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
		Addr:         srv.port,
		Handler:      srv,
		ReadTimeout:  1 * time.Second,  // in case of a too long request?
		WriteTimeout: 10 * time.Second, // in case of no TCP-ACK
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
