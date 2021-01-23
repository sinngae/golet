package main

import (
	"fmt"

	"github.com/sinngae/goerrcode"
	"github.com/sinngae/goerrcode/sample/errcode"
)

func main() {
	fmt.Println("hi")

	err := goerrcode.WithRetCodeMessage(fmt.Errorf("this is caust"), errcode.DBNoData)
	err = goerrcode.WithStack(err)
	err = goerrcode.WithRetCodeMessage(err, errcode.NoParam)
	err = goerrcode.WithRetCode(err, errcode.RetCodeParamInvalid)
	err = goerrcode.WithRetCodeMessagef(err, 404, "this is msg fmt, %s", "this is msg")

	fmt.Println(err)
	fmt.Println(goerrcode.Message(err))
	fmt.Println(goerrcode.RetCode(err))
	fmt.Println(goerrcode.RetCodeMsg(err))
	fmt.Println(goerrcode.Stack(err))

	fmt.Println("work")
}
