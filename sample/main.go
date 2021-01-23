package main

import (
	"fmt"

	"github.com/sinngae/goerrcode"
	"github.com/sinngae/goerrcode/sample/errcode"
)

func main() {
	fmt.Println("hi, work")
	err := goerrcode.WithRetCodeMessage(fmt.Errorf("test"), errcode.DBNoData)
	err = goerrcode.WithStack(err)
	fmt.Println(err)
	msg := goerrcode.Message(err)
	fmt.Println(msg)
	fmt.Println("hi")
}
