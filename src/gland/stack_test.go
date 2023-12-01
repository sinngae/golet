package gland

import (
	"fmt"
)

func ExampleStack() {
	st := stack(0)
	str := string(st)
	fmt.Println(str)
	// output:
	// afsdf
}

func ExampleStack1() {
	st := Stack()
	str := string(st)
	fmt.Println(str)
	// output:

}
