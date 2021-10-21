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
