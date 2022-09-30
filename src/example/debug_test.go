package main

import (
	"log"

	"github.com/sinngae/golet/src/debug"
)

func main() {
	if debug.Debug {
		log.Printf("service run under debug mode...\n")
	}
	log.Printf("service run...")
	// output

}
