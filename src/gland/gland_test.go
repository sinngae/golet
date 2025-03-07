package gland

import (
	"fmt"
)

var (
	err    = fmt.Errorf("this is a error")
	gland0 = New(404, "this is msg", err)
	gland1 = NewWithStack(404, "this is msg", err)
	gland2 = gland0.WithStack()
)

func ExampleGlandOutput() {
	fmt.Println(err)
	fmt.Println(gland0)
	fmt.Println(gland1)
	fmt.Println(gland2)

	// output:

}
