package stdst

import (
	"fmt"
)

func ExampleSliceRange() {
	data := []int{0, 2, 4, 6, 3, 4}
	s1 := data[:100]
	s2 := data[3:]
	println(s1, s2)
	//output:

}

func SliceMod(list []int, val int) {
	fmt.Printf("%p %v %+v %#v\n", list, list, list, list)
	list[0] = val
}

func SliceAppend(list []int, val int) {
	list[0] = 999
	fmt.Printf("%p %v %+v %#v\n", list, list, list, list)
	list = append(list, 1)
	fmt.Printf("%p %v %+v %#v\n", list, list, list, list)
	list[1] = 111
}

func ExampleSlice() {
	data := []int{0}
	SliceMod(data, 1)
	fmt.Printf("%p %v %+v %#v\n", data, data, data, data)
	SliceAppend(data, 2)
	fmt.Printf("%p %v %+v %#v\n", data, data, data, data)
	//output:

}

func ExampleSliceCompare() {
	s1 := struct {
		_   int
		val int
	}{val: 0}
	s2 := struct{ val int }{0}
	fmt.Println( /*s1 == s2*/ s1, s2)
	//output:
}
