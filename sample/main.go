package main

import (
	"fmt"

	"github.com/sinngae/gland"
	"github.com/sinngae/gland/sample/errcode"
)

func main() {
	fmt.Println("hi")

	err := gland.WithRetCodeMessage(fmt.Errorf("this is caust"), errcode.DBNoData)
	err = gland.WithStack(err)
	err = gland.WithRetCodeMessage(err, errcode.NoParam)
	err = gland.WithRetCode(err, errcode.RetCodeParamInvalid)
	err = gland.WithRetCodeMessagef(err, 404, "this is msg fmt, %s", "this is msg")

	fmt.Println(err)
	fmt.Println(gland.Message(err))
	fmt.Println(gland.RetCode(err))
	fmt.Println(gland.RetCodeMsg(err))
	fmt.Println(gland.Stack(err))

	fmt.Println("work")
}
