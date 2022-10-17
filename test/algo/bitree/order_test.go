package bitree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	BiTreeNode struct {
		Val   int
		Left  *BiTreeNode
		Right *BiTreeNode
	}
)

func preorderTraversal(root *BiTreeNode) []int {
	if root == nil {
		return nil
	}
	dest := []int{root.Val}
	left := preorderTraversal(root.Left)
	dest = append(dest, left...)
	right := preorderTraversal(root.Right)
	dest = append(dest, right...)
	return dest
}

func sufOrderTraversal(root *BiTreeNode) []int {
	if root == nil {
		return nil
	}
	left := sufOrderTraversal(root.Left)
	right := sufOrderTraversal(root.Right)
	dest := left
	dest = append(dest, right...)
	dest = append(dest, root.Val)
	return dest
}

func midOrderTraversal(root *BiTreeNode) []int {
	if root == nil {
		return nil
	}
	left := midOrderTraversal(root.Left)
	right := midOrderTraversal(root.Right)
	dest := left
	dest = append(dest, root.Val)
	dest = append(dest, right...)
	return dest
}

func TestPreOrderTraversal(t *testing.T) {
	tts := []struct {
		name string
		data *BiTreeNode
		want []int
	}{
		{
			name: "case 0",
			data: &BiTreeNode{Val: 1, Left: &BiTreeNode{Val: 2}, Right: &BiTreeNode{Val: 3}},
			want: nil,
		},
	}

	for _, tt := range tts {
		got := preorderTraversal(tt.data)
		assert.Equal(t, tt.want, got)
		got1 := sufOrderTraversal(tt.data)
		assert.Equal(t, tt.want, got1)
		got2 := sufOrderTraversal(tt.data)
		assert.Equal(t, tt.want, got2)
	}
}
