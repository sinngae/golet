package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sinngae/golet/cmd/testcase/gin/internal/route"
)

func main() {
	defer func() {
		if rcv := recover(); rcv != nil {
			fmt.Printf("serivce panic, recover=%v", rcv)
		}
	}()

	startAt := time.Now()
	fmt.Printf("service start @ %s\n", startAt)

	err := bootstrap()
	if err != nil {
		fmt.Printf("bootstrap failed, err=%v\n", err)
		return
	}

	for {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		sig := <-c
		log.Printf("Syscall: %v", sig)
	}
}

func bootstrap() error {
	service := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: ginInit(),
	}

	go func() {
		fmt.Printf("service listen to %v\n", 8080)
		if err := service.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Listen failed, err=%v\n", err)
		}
		fmt.Printf("service quit @ %v\n", time.Now())
	}()

	return nil
}

func ginInit() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	service := gin.New()
	// service.Use(middleware)
	route.Init(service)
	return service
}
