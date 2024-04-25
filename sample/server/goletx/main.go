package main

import (
	"github.com/sinngae/golet/sample/server/internal/gin"
)

func main() {
	startParam := "port=:8080"
	gin.Startup(startParam)
}
