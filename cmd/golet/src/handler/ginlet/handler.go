package ginlet

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Handle() error {
	startAt := time.Now()
	fmt.Printf("service start @ %s\n", startAt)

	err := startup()
	if err != nil {
		fmt.Printf("startup failed, err=%s\n", err)
		return err
	}

	return nil
}

func signalWatch() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		for sig := range c {
			log.Printf("Syscall: %v", sig)
		}
	}()
}

func startup() error {
	service := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: NewService(),
	}

	fmt.Printf("service listen to %v\n", 8080)
	if err := service.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Listen failed, err=%v\n", err)
		return err
	}
	fmt.Printf("service quit @ %v\n", time.Now())

	return nil
}
