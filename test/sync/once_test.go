package sync

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	data string
)

func ExampleOnce() {
	once.Do(func() {
		data = "hi,once"
	})
	fmt.Println(data)
	// output:
	// "hi,once"
}

type OnceSt struct {
	val  int
	once sync.Once
}

func (st *OnceSt) GetVal() int {
	once.Do(func() {
		st.val++
	})
	return st.val
}

func ExampleOnceSt() {
	var demo OnceSt
	fmt.Println(demo.GetVal())
	fmt.Println(demo.GetVal())
	fmt.Println(demo.GetVal())
	// output:
	// 1
	// 1
	// 1
}
