package init

import (
	"fmt"

	a2 "github.com/sinngae/golet/test/init/a"
	common2 "github.com/sinngae/golet/test/init/common"
)

func ExampleInit() {
	aVal := a2.ValA
	println(aVal)
	val := common2.Val
	fmt.Println(val)
	println("hi, this")
	//output:
	// 10
}
