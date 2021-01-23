package goerrcode

import (
	"fmt"
	"testing"
)

var errWithRC = WithRetCode(fmt.Errorf("this is cause"), 404)

func TestWithRetCode(t *testing.T) {
	fmt.Println(errWithRC)
}

func TestRetCode(t *testing.T) {
	fmt.Println(RetCode(errWithRC))
}
