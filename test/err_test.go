package test

import (
	"fmt"

	"github.com/pkg/errors"
)

func ExampleErrPrint() {
	err := errors.New("this bad err")
	err = errors.Wrapf(err, "this fmt")
	msg := fmt.Sprintf("err=%v", err)
	println(msg)
	msg = fmt.Sprintf("err=%s", err)
	println(msg)
	msg = fmt.Sprintf("err=%d", err)
	println(msg)
	err = fmt.Errorf("err=%w", err)
	msg = fmt.Sprintf("err=%s", err)
	println(msg)
	// output:

}
