package test

import "fmt"

func ExampleRecover() {
	ss("xx", func() {
		//defer cronRecover()()
		defer abc()
		fmt.Println("this is pnc")
		panic("this")
	})
	fmt.Println("Hello world!")
	// output:

}

func abc() {
	recv()
}

func recv() {
	fmt.Println("this is recv")
	if err := recover(); err != nil {
		fmt.Println(err, "abc def")
	}
}

func cronRecover() func() {
	fmt.Println("this is out")
	return func() {
		if err := recover(); err != nil {
			fmt.Println(err, "abc def")
		}
	}
}
func ss(s string, f func()) {
	f()
	fmt.Println(s)
}
