package gin

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Startup(param string) {
	startAt := time.Now()
	fmt.Printf("service starting... @ %s\n", startAt)

	var srv *Server
	go func() {
		sp := parseStartParam(param)
		srv = NewServer(sp.port)
		err := srv.Startup()
		if err != nil {
			log.Fatalf("startup err:%s", err.Error())
		}
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

type StartParam struct {
	port string
}

func parseStartParam(param string) *StartParam {
	valList, err := url.ParseQuery(param)
	if err != nil {
		panic(fmt.Sprintf("parse %s failed, err=%s", param, err))
		return nil
	}

	sp := &StartParam{}
	for k, v := range valList {
		switch k {
		case "port":
			sp.port = v[0]
		}
	}
	return sp
}
