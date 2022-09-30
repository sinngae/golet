package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	ng := gin.Default()
	ng = Route(ng)

	//err := http.ListenAndServe("127.0.0.1:8080", ng)
	err := ng.Run() // default port: 8080
	if err != nil {
		log.Fatalf("gin Run failed, err=%s\n", err)
	}
}
