package go_stl

import (
	"fmt"
	"strconv"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type (
	isNilFn func(data interface{}) bool
	demoSt  struct {
		val int
	}
	demoFn func()
)

var (
	badIsNil = func(data interface{}) bool {
		return data == nil
	}
	ptrSt  = (*demoSt)(nil)
	ptrSt1 *demoSt
	ptrInt *int
	intV   int

	ptrFn = (*demoFn)(nil)
	chanV chan int
	itfV  interface{}
)

func ExampleIsNil() {
	isNil := IsNilW
	fmt.Printf("nil: rawRes:omit, badRes:%t, res:%t\n" /*nil==nil*/, badIsNil(nil), isNil(nil))
	fmt.Printf("int const: rawRes:omit, badRes:%t, res:%t\n" /*1 == nil*/, badIsNil(1), isNil(1))
	fmt.Printf("struct ptr: rawRes:%t, badRes:%t, res:%t\n", ptrSt == nil, badIsNil(ptrSt), isNil(ptrSt))
	fmt.Printf("struct ptr2: rawRes:%t, badRes:%t, res:%t\n", ptrSt1 == nil, badIsNil(ptrSt1), isNil(ptrSt1))
	fmt.Printf("int val: rawRes:omit, badRes:%t, res:%t\n" /*intV == nil, */, badIsNil(intV), isNil(intV))
	fmt.Printf("int ptr: rawRes:%t, badRes:%t, res:%t\n", ptrInt == nil, badIsNil(ptrInt), isNil(ptrInt))
	fmt.Printf("func ptr: rawRes:%t, badRes:%t, res:%t\n", ptrFn == nil, badIsNil(ptrFn), isNil(ptrFn))
	fmt.Printf("chan val: rawRes:%t, badRes:%t, res:%t\n", chanV == nil, badIsNil(chanV), isNil(chanV))
	fmt.Printf("interface val: rawRes:%t, badRes:%t, res:%t\n", itfV == nil, badIsNil(chanV), isNil(chanV))
	// output:

}

func TestIsNilCase0(t *testing.T) {
	t.Run("-1", func(t *testing.T) {
		got := badIsNil(nil)
		//assert.Equal(t, nil == nil, got) // invalid operation: nil == nil (operator == not defined on nil)
		assert.Equal(t, true, got)
		got = IsNil(nil)
		assert.Equal(t, true, got)
	})
	t.Run("-2", func(t *testing.T) {
		var a *int
		got := badIsNil(a)
		assert.NotEqual(t, a == nil, got)
		got = IsNil(a)
		assert.Equal(t, a == nil, got)
	})
	t.Run("-3", func(t *testing.T) {
		var a *demoSt
		got := badIsNil(a)
		assert.NotEqual(t, a == nil, got)
		got = IsNil(a)
		assert.Equal(t, a == nil, got)
	})
	t.Run("-4", func(t *testing.T) {
		var a *demoFn
		got := badIsNil(a)
		assert.NotEqual(t, a == nil, got)
		got = IsNil(a)
		assert.Equal(t, a == nil, got)
	})
}

func TestIsNil(t *testing.T) {
	tests := []interface{}{
		nil, 1, "abc",
		(*int)(nil),
		(*demoSt)(nil),
		(map[string]int)(nil),
		([]int)(nil),
		(chan int)(nil),
		(*demoFn)(nil),
		unsafe.Pointer(nil),
		(interface{})(nil),
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := IsNil(tt)
			assert.Equal(t, tt == nil, got)
		})
	}
}
