package goerrcode

import (
	"fmt"
	"testing"
)

var (
	errCause      = fmt.Errorf("this is cause")
	rcMsg         = NewRetCodeMsg(404, "this is msg")
	errWithRcMsg  = WithRetCodeMessage(errCause, rcMsg)
	errWithRcMsgf = WithRetCodeMessagef(errCause, 404, "this is msg fmt, %s", "this is msg")
)

func TestWithRetCodeMessage(t *testing.T) {
	fmt.Println(errWithRcMsg)
}

func TestWithRetCodeMessagef(t *testing.T) {
	fmt.Println(errWithRcMsgf)
}

func TestRetCodeMsg(t *testing.T) {
	fmt.Println(RetCodeMsg(errWithRcMsg))
}

func TestRcMsg(t *testing.T) {
	fmt.Println(rcMsg)
}
