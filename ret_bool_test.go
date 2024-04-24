package main

func main() {
	ok, val := fRet()
	println(ok, val)
}

func fRet() (bool, int) {
	cnt := 0
	return cnt > 0, setCnt(&cnt)
}

func setCnt(val *int) int {
	*val = 10
	return 2
}
