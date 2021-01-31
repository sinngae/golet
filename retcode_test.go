package gland

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

func TestIs(t *testing.T) {
	fmt.Println(Is(errWithRC, 404))
	fmt.Println(Is(errWithRC, 403))
}
