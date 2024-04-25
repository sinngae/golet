package go_stl

import (
	"fmt"
	"reflect"
	"unsafe"
)

// check whether data is nil, INNER function receive data as interface{}

// IsNil
// The expression (*[2]uintptr)(unsafe.Pointer(&s.i))[1], works, but there's no guarantee that it will work in the future.
// The Value.Pointer method is the supported way to solve this problem. Other solutions require that you make assumptions
// about the memory layout.
func IsNil(data interface{}) bool {
	return (*[2]uintptr)(unsafe.Pointer(&data))[1] == 0
}

func IsNilW(data interface{}) bool {
	if data == nil {
		return true
	}
	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.Array, reflect.String, reflect.Struct:
		return false
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return val.IsNil()
	default:
		// This should never happens, but will act as a safeguard for
		// later, as a default value doesn't makes sense here.
		panic(fmt.Errorf("unsupported kind:%s, IsNil", val.Kind()))
	}
}
