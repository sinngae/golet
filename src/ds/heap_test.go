package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLHeap_Push(t *testing.T) {
	tests := []struct {
		name  string
		list  []int
		want  []int
		want2 []int
	}{
		{
			name:  "case 0",
			list:  []int{1, 5, -1, 2, -1, 9, -1, -1, 3, 22, 7, 10, 33, 100, 21, 4},
			want:  []int{},
			want2: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := LHeap{}
			for _, item := range tt.list {
				heap.Push(item)
			}
			assert.Equal(t, tt.want, heap.Data)
			got := make([]int, 0)
			for {
				val, ok := heap.Pop()
				if !ok {
					break
				}
				got = append(got, val)
			}
			assert.Equal(t, tt.want2, got)
		})
	}
}

func TestRHeap_Push(t *testing.T) {
	tests := []struct {
		name  string
		list  []int
		want  []int
		want2 []int
	}{
		{
			name:  "case 0",
			list:  []int{1, 5, 2, 9, 3, 22, 7, 10, 33, 100, 21, 4},
			want:  []int{},
			want2: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := RHeap{}
			for _, item := range tt.list {
				heap.Push(item)
			}
			assert.Equal(t, tt.want, heap.Data)
			got := make([]int, 0)
			for {
				val, ok := heap.Pop()
				if !ok {
					break
				}
				got = append(got, val)
			}
			assert.Equal(t, tt.want2, got)
		})
	}
}
