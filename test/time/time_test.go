package time

import (
	"fmt"
	"time"
)

func ExampleTime() {
	val := time.Now().UnixNano()
	fmt.Println(val)
	// output:

}

func ExampleTimer() {
	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	fmt.Printf("time end\n")
	//output:
}

func ExampleTick() {
	tick := time.NewTicker(3 * time.Second)
	for range tick.C {
		fmt.Printf("")
	}
}
