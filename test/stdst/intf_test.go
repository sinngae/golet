package stdst

func ExampleInterface() {
	var itf1, itf2 interface{}
	itf1 = nil
	println(itf1 == nil)
	itf2 = 2
	itf2 = "nil"
	println(itf2 == nil)

	ptr := (*Demo)(nil)
	println(itf1 == ptr)

	println(itf1 == itf2)
	itf1, itf2 = "123", 123
	println(itf1 == itf2)
	//output:

}

// 编译报错
//# github.com/sinngae/golet/test/stdst [github.com/sinngae/golet/test/stdst.test]
//.\intf_test.go:31:8: cannot use type ItfA outside a type constraint: interface contains type constraints
//type (
//	StA struct {
//		//Val int
//	}
//	ItfA interface {
//		StA
//	}
//)
//
//func ExampleFunA() {
//	var a ItfA
//	//a.Val = 1
//	println(a)
//	//output:
//
//}
