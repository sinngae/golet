package main

import (
	"github.com/sinngae/gland/pkg/debug"
	"log"
)

func main() {
	if debug.IsDebugging(false) {
		log.Printf("service run under debug mode...\n")
	}
	log.Printf("service run...")
	// output
	
}
