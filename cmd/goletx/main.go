package main

import (
	"github.com/sinngae/golet/cmd/internal/gin"
)

func main() {
	startParam := "port=:8080"
	gin.Startup(startParam)
}
