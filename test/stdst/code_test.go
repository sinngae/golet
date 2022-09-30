package stdst

func ExampleSwap() {
	valA, valB := 1, 2
	valA, valB = valB, valA

	println(valA, valB)
	// output:
	//2 1
}

func ExampleDefer() {
	//debug.SetMaxThreads()
	fn := func() int {
		i := 0
		defer func() {
			println(i)
			i++
			println(i)
		}()
		return i
	}
	val := fn()
	println(val)
	//output:
	// 0
	// 1
	// 0
}

func ExampleUint() {
	a, b := uint(1), uint(2)
	c := a - b
	println(c)
	//output:
}
