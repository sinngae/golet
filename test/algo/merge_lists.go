package algo

import (
	"fmt"
	"strconv"
	"strings"
)

//import . "nc_tools"
/*
* type ListNode struct{
*   Val int
*   Next *ListNode
* }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	str := strings.Builder{}
	iter := node
	for {
		if iter == nil {
			break
		}
		str.WriteString(strconv.Itoa(iter.Val))
		str.WriteString(",")
		iter = iter.Next
	}
	return str.String()
}

/**
 *
 * @param lists ListNode类一维数组
 * @return ListNode类
 */
func mergeKLists(lists []*ListNode) *ListNode {
	// write code here
	if len(lists) <= 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	heap := ListHeap{}
	for _, item := range lists {
		if item == nil {
			continue
		}
		heap.Push(item)
	}

	dest := &ListNode{}
	dTmp := dest
	for {
		temp := heap.Pop()
		val, ok := heap.First()
		if !ok {
			dTmp.Next = temp
			break
		}
		for {
			if temp == nil {
				break
			}
			if temp.Val > val {
				break
			}
			dTmp.Next = &ListNode{Val: temp.Val}
			dTmp = dTmp.Next
			temp = temp.Next
		}
		if temp != nil {
			heap.Push(temp)
		}
		str1, str2 := dest.String(), temp.String()
		fmt.Sprintf("dest: %s, temp: %s\n", str1, str2)
	}
	return dest.Next
}

type (
	ListHeap struct {
		Data []*ListNode
	}
)

func (heap *ListHeap) First() (int, bool) {
	if len(heap.Data) <= 0 {
		return 0, false
	}
	return heap.Data[0].Val, true
}

func (heap *ListHeap) Push(node *ListNode) {
	if node == nil {
		return
	}
	heap.Data = append(heap.Data, node)
	idx := len(heap.Data) - 1
	heap.reorder(idx)
}

func (heap *ListHeap) reorder(idx int) {
	for {
		if idx == 0 {
			return
		}

		next := (idx - 1) / 2
		if heap.Data[next].Val <= heap.Data[idx].Val {
			return
		}
		heap.Data[next], heap.Data[idx] = heap.Data[idx], heap.Data[next]
		idx = next
	}
}

func (heap *ListHeap) Pop() *ListNode {
	if len(heap.Data) <= 0 {
		return nil
	}
	if len(heap.Data) == 1 {
		val := heap.Data[0]
		heap.Data = heap.Data[:0]
		return val
	}

	val := heap.Data[0]

	max := len(heap.Data) - 1
	idx := 0
	for {
		next := idx*2 + 1
		if next >= max {
			heap.Data[idx] = heap.Data[max]
			heap.reorder(idx)
			break
		}
		if heap.Data[next].Val >= heap.Data[next+1].Val {
			heap.Data[idx] = heap.Data[next+1]
			idx = next + 1
			continue
		}
		heap.Data[idx] = heap.Data[next]
		idx = next
	}
	heap.Data = heap.Data[:max]
	return val
}
