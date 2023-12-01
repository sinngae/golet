package bitree

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param root TreeNode类
 * @return TreeNode类
 */
func Serialize(root *TreeNode) string {
	// write code here
	if root == nil {
		return "#"
	}
	left := Serialize(root.Left)
	right := Serialize(root.Right)
	return fmt.Sprintf("%d,%s,%s", root.Val, left, right)
}
func Deserialize(s string) *TreeNode {
	// write code here
	list := strings.Split(s, ",")
	if len(list) <= 0 {
		return nil
	}

	root := toNode(list[0])
	data := []*TreeNode{root}
	idx := 1
	preNode := root
	for {
		leftChild := toNode(list[idx])
		preNode.Left = leftChild
		if leftChild != nil {
			data = append(data, leftChild)
			preNode = leftChild
		} else {
			//preNode
		}
		idx++
		if idx >= len(list) {
			break
		}
		rightChild := toNode(list[idx])
		//parent.Right = rightChild
		if rightChild != nil {
			data = append(data, rightChild)
		}
		idx++
		if len(data) <= 0 {
			break
		}
	}
	return root
}

func toNode(str string) *TreeNode {
	if str == "#" {
		return nil
	}
	val, _ := strconv.Atoi(str)
	return &TreeNode{Val: val}
}

func TestSerialize(t *testing.T) {
	tcs := []struct {
		str  string
		want string
	}{
		{
			str:  "1,2,3,4,5,6,#,7,#,#,8,#",
			want: "1,2,3,4,5,6,#,7,#,#,8,#",
		},
	}
	for _, tc := range tcs {
		t.Run("case 0", func(t *testing.T) {
			node := Deserialize(tc.str)
			got := Serialize(node)
			fmt.Printf("str:%s\n", got)
			assert.Equal(t, tc.want, got)
		})
	}
}
