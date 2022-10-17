package abc

func ExampleAbc() {
	val := swapContents([]int{1, 2, 3, 4})
	println(val)
	//output:

}
func swapContents(arrData []int) []int {
	// write code here
	mid := (len(arrData) - 1) / 2
	for idx := 0; idx < mid; idx++ {
		idx2 := len(arrData) - idx
		arrData[idx], arrData[idx2] = arrData[idx2], arrData[idx]
	}
	return arrData
}
