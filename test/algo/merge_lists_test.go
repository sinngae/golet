package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func trans2Chain(list []int) *ListNode {
	var dest *ListNode
	iter := &dest
	for _, item := range list {
		*iter = &ListNode{Val: item}
		iter = &((*iter).Next)
	}
	return dest
}

func trans2KLists(list [][]int) []*ListNode {
	dest := make([]*ListNode, 0, len(list))
	for _, item := range list {
		node := trans2Chain(item)
		dest = append(dest, node)
	}
	return dest
}

func trans2List(klist *ListNode) []int {
	if klist == nil {
		return nil
	}
	dest := make([]int, 0)
	iter := klist
	for {
		val := iter.Val
		dest = append(dest, val)
		iter = iter.Next
		if iter == nil {
			break
		}
	}
	return dest
}

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  *ListNode
	}{
		{
			name: "abc",
			lists: [][]int{{-7, -7, -6, -4, -4, 1, 1}, {-10, -9, -8, -7, -7, -5, -2, 2}, {-10, 3},
				{}, {}, {-5, -4, -4, -4, -1, -1, 4}, {}, {-10, -5, -3, -1, 1, 2}, {-10, -9, -5, -4, -3, -1, -1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeKLists(trans2KLists(tt.lists))
			got1 := trans2List(got)
			assert.Equalf(t, tt.want, got1, "mergeKLists(%v)", tt.lists)
		})
	}
}
