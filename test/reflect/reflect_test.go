package reflect

import (
	"fmt"
	"reflect"
)

var (
	intCh chan int
	fn    = func() {}
)

func ExampleReflect() {
	tests := []interface{}{
		nil,
		"abc",
		123,
		struct {
			val int
		}{12},
		intCh,
		fn,
		[]string{"def", "asf"},
		map[string]int{"sdf": 1},
	}
	for _, tt := range tests {
		fmt.Printf("type:%v, val:%v\n", reflect.TypeOf(tt), reflect.ValueOf(tt))
		reflect.Value{}.Interface()
	}
	// output:

}
