package reverse_proxy

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartUp() {
	startAt := time.Now()
	fmt.Printf("service starting... @ %s\n", startAt)

	var srv *Server
	go func() {
		srv = NewServer(":8080")
		_ = srv.Startup()
	}()

	// exit elegantly
	for {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		sig := <-c
		log.Printf("server shutting down... signal=%v", sig)
		break
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
