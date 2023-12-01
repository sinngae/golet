package go_std

import (
	"fmt"
	"testing"
)

func testFunc(data map[string]int) {
	data["123"] = 1
}

func TestMapDeliver(t *testing.T) {
	data := make(map[string]int)
	testFunc(data)
	fmt.Println(data)
}
