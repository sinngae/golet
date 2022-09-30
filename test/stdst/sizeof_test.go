package stdst

import (
	"fmt"
	"reflect"
)

type (
	Demo struct {
		v0 byte
		v1 int
		v2 float32
		v3 string
		v4 []int
		v5 map[string]string
		v6 chan int
		v7 func() int
		Base
		BFunc
	}
	Base struct {
		v0 int
	}
	BFunc func() int
)

func ExampleSizeOf() {
	typ := reflect.TypeOf(Demo{})
	fmt.Printf("St is %d bytes long\n", typ.Size())

	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d, tag=%s, pkgpath=%s, idx=%d, anony=%t\n",
			field.Name, field.Offset, field.Type.Size(), field.Type.Align(),
			field.Tag, field.PkgPath, field.Index, field.Anonymous)
	}
	// output:
	//
}
