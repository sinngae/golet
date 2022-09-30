package reflect

import (
	"fmt"
	"reflect"
	"unsafe"
)

type MySt struct {
	byteV   byte
	shortV  int16
	longVal int32
	//intVal   int
	sliceVal []byte
}

func ExampleMySt() {
	typ := reflect.TypeOf(MySt{})
	fmt.Printf("Struct is %d bytes long\n", typ.Size())
	// We can run through the fields in the structure in order
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
	// output:
	//

}

func ExampleSliceHeader() {
	typ := reflect.TypeOf(reflect.SliceHeader{})
	fmt.Printf("Struct is %d bytes long\n", typ.Size())
	// We can run through the fields in the structure in order
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
	//reflect.Map
	// output:
	//

}

func ExampleMySt2() {
	data := MySt{
		byteV:    0x1,
		shortV:   0x0203,
		longVal:  0x04050607,
		sliceVal: []byte{0x08, 0x09, 0x0a},
	}
	dataBytes := (*[32]byte)(unsafe.Pointer(&data))
	fmt.Printf("Bytes are %#v\n", dataBytes)
	dataSlice := *(*reflect.SliceHeader)(unsafe.Pointer(&data.sliceVal))
	fmt.Printf("Slice data is %#v\n", (*[3]byte)(unsafe.Pointer(dataSlice.Data)))
	// output:
	// Bytes are &[32]uint8{0x1, 0x0, 0x3, 0x2, 0x7, 0x6, 0x5, 0x4, 0xd0, 0x82, 0x1, 0x0, 0xc0, 0x0, 0x0, 0x0, 0x3,
	// 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	//
	// Slice data is &[3]uint8{0x8, 0x9, 0xa}
}
