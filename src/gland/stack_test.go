<<<<<<< HEAD
package gland

import (
	"fmt"
	"testing"
)

var (
	errWithStack = WithStack(fmt.Errorf("this is cause"))
)

func TestWithStack(t *testing.T) {
	fmt.Println(errWithStack)
}

func TestNewStack(t *testing.T) {
	fmt.Println(NewStack())
}

func TestStack(t *testing.T) {
	fmt.Printf("%+v", Stack(errWithStack))
}
=======
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
>>>>>>> 4aad3e1b64427d5ebafb07f037b140f7eb3a6511
