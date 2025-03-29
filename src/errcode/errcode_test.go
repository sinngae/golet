package errcode

import (
	"fmt"
)

func ExampleErrcode() {
	fmt.Println("hi")

	err2 := New(RetCodeParamInvalid, "")

	fmt.Println(DBNoData)
	fmt.Println(err2)

	fmt.Println("work")
	//output:

}
