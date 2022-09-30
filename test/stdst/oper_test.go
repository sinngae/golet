package stdst

func ExampleAddqual() {
	a, b := 1, 2
	//a, b += 1, 2 // not suport, 意外的 +=，应为 ':='、'=' 或 ','
	println(a, b, a^b)
	//output:

}
