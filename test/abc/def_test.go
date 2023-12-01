package abc

import (
	"fmt"
	"testing"
)

type Node struct {
	data interface{}
	next *Node
}

func TestFn(t *testing.T) {
	first := &Node{data: 1}
	second := &Node{data: 2}
	third := &Node{data: 3}

	first.next = second
	second.next = third

	val := Fn(first)
	fmt.Println(val)
}

func Fn(src *Node) *Node {
	if src == nil {
		return nil
	}

	dest := &Node{
		data: src.data,
		next: nil,
	}
	temp := src.next
	for {
		if temp == nil {
			break
		}
		dest = &Node{
			data: temp.data,
			next: dest,
		}
		temp = temp.next
	}
	return dest
}
