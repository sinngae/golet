package a

import "github.com/sinngae/gland/bench/init/common"

var ValA int

func init() {
	common.Val = 10
	ValA = 11
	println("hi, a")
}
