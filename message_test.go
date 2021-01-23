package goerrcode

import (
	"fmt"
	"testing"
)

var errWithMsg = WithMessage(fmt.Errorf("this is cause"), "this is msg")

func TestWithMessage(t *testing.T) {
	fmt.Println(errWithMsg)
}

func TestMessage(t *testing.T) {
	fmt.Println(Message(errWithMsg))
}
