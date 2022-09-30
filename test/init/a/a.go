package a

import (
	common2 "github.com/sinngae/golet/test/init/common"
)

var ValA int

func init() {
	common2.Val = 10
	ValA = 11
	println("hi, a")
}
