package assert

import "reflect"

func Equal(src, dst interface{}) {
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)
	if srcType != dstType {
		panic("type not equal")
	}

}
