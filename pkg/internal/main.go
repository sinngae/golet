package main

import (
	"log"

	"github.com/sinngae/gland/pkg/internal/debug_"
)

func main() {
	if debug_.IsDebugging(false) {
		log.Printf("service run under debug mode...\n")
	}
	log.Printf("service run...")
}
