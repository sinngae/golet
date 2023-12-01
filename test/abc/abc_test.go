package abc

import "fmt"

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

func ExampleMStr() {
	//val := mStr("abcabcabcd", "abcd")
	val := mStr("abchabccabccd", "abcd")
	fmt.Printf("val=%d\n", val)
	//output:
}

func mStr(str string, mStr string) int {
	list := []byte(str)
	mList := []byte(mStr)
	if len(list) < len(mList) {
		return -1
	}

	stats := make(map[byte]struct{})
	for _, ch := range mList {
		stats[ch] = struct{}{}
	}

	tMap := make(map[byte]int)
	for _, ch := range list[:len(mList)] {
		if _, ok := stats[ch]; ok {
			tMap[ch] = tMap[ch] + 1
		}
	}
	if len(tMap) == len(mList) {
		return 0
	}

	idx := 0
	for _, ch := range list[len(mList):] {
		if val, ok := tMap[list[idx]]; ok {
			if val == 1 {
				delete(tMap, list[idx])
			} else {
				tMap[list[idx]] = val - 1
			}
		}
		if _, ok := stats[ch]; ok {
			tMap[ch] += 1
		}
		if len(tMap) == len(mList) {
			return idx + 1
		}
		idx++
	}
	return -1
}
