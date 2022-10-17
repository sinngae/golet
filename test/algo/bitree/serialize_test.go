package bitree

import (
	"fmt"
	"strconv"
	"strings"
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

	root := &TreeNode{}
	data := []*TreeNode{toNode(list[0])}
	idx := 1
	for {
		node := data[0]
		data = data[1:]

	}
}

func toNode(str string) *TreeNode {
	if str == "#" {
		return nil
	}
	val, _ := strconv.Atoi(str)
	return &TreeNode{Val: val}
}
