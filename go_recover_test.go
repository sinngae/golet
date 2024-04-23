package main

import "fmt"

func main() {
	ss("xx", func() {
		defer cronRecover()()
		fmt.Println("this is pnc")
		panic("this")
	})
	fmt.Println("Hello world!")
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
