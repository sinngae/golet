package errcode

import (
	"fmt"

	"github.com/sinngae/golet/src/gland"
)

func ExampleErrcode() {
	fmt.Println("hi")

	err := DBNoData
	err0 := err.WithStack()
	err1 := NoParam.WithCause(fmt.Errorf("this is a err"))
	err2 := gland.New(RetCodeParamInvalid, "", err)
	err3 := gland.New(404, "this is msg fmt, %s", err)

	fmt.Println(DBNoData)
	fmt.Println(err0)
	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)

	fmt.Println("work")
	//output:

}
