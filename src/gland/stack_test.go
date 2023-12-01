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
