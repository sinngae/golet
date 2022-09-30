package ants

import (
	"log"

	"github.com/panjf2000/ants/v2"
)

func ExampleAnts() {
	pool, err := ants.NewPool(10)
	if err != nil {
		log.Fatalf("new pool fail, err=%s", err)
	}

	pool.Tune()
}
