package reverse_proxy

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	port string

	entity *ReverseProxy
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (srv *Server) Startup() error {
	rp := NewReverseProxy(nil) // todo

	server := &http.Server{
		Addr:         srv.port,
		ReadTimeout:  1 * time.Second,  // in case of a too long request?
		WriteTimeout: 10 * time.Second, // in case of no TCP-ACK
		Handler:      rp,
	}
	srv.entity = rp

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
	srv.entity.Shutdown()
	//if err := srv.entity.Shutdown(ctx); err != nil {
	//	fmt.Printf("server shutdown failed, err=%s", err)
	//	return err
	//}
	return nil
}
