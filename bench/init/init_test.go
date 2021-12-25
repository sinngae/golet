package init

import (
	"fmt"

	"github.com/sinngae/gland/bench/init/a"
	"github.com/sinngae/gland/bench/init/common"
)

func ExampleInit() {
	aVal := a.ValA
	println(aVal)
	val := common.Val
	fmt.Println(val)
	println("hi, this")
	//output:
	// 10
}
