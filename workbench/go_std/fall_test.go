package go_std

import "fmt"

func ExampleFallThrough() {
	switch 1 {
	case 0:
		println(0)
		fallthrough
	case 1:
		println(1)
		fallthrough
	case 2:
		println(2)
	case 3:
		println(3)
	}
	//output:
}

func ExampleList() {
	list := []string{"abc", "dafa", "dafa"}
	str := fmt.Sprintf("%s", list)
	str2 := fmt.Sprintf("%v", list)
	println(str, str2)
	// output:
	//
}
