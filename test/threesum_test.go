package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func threeSum(nums []int) [][]int {
	srcMap := make(map[string]int)

	dest := make([][]int, 0)
	max := len(nums)
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			if j == i {
				continue
			}

			for k := 0; k < max; k++ {
				if k == i || k == j {
					continue
				}

				if nums[i]+nums[j]+nums[k] == 0 {
					list := sortList(nums[i], nums[j], nums[k])
					//slice.sort

					str := fmt.Sprintf("%d#%d#%d", list[0], list[1], list[2])
					if _, ok := srcMap[str]; ok {
						continue
					} else {
						srcMap[str] = 0
					}
					dest = append(dest, list)
				}
			}
		}
	}
	return dest
}

func sortList(num1, num2, num3 int) []int {
	if num1 > num2 {
		if num1 > num3 {
			if num2 > num3 {
				return []int{num3, num2, num1}
			} else {
				return []int{num2, num3, num1}
			}
		} else {
			return []int{num2, num1, num3}
		}
	} else { // num1 < num2
		if num1 < num3 {
			if num2 < num3 {
				return []int{num1, num2, num3}
			} else {
				return []int{num1, num3, num2}
			}
		} else {
			return []int{num3, num1, num2}
		}
	}
}

func TestThreeSum(t *testing.T) {
	//data := []int{-1, 0, 1, 2, -1, -4}
}

func sortInts(src []int) []int {
	for j := 0; j < len(src)-1; j++ {
		max := len(src) - j - 1
		for i := 0; i < max; i++ {
			if src[i] > src[i+1] {
				temp := src[i]
				src[i] = src[i+1]
				src[i+1] = temp
			}
		}
	}
	return src
}

func TestSortInts(t *testing.T) {
	tcs := []struct {
		tc   string
		src  []int
		want []int
	}{
		{tc: "case 0", src: []int{0, 1, -1}, want: []int{-1, 0, 1}},
		{tc: "sadf", src: []int{34, 324, 3, -5, -231, 546, 23}, want: []int{0, 0}},
	}

	for _, tc := range tcs {
		got := sortInts(tc.src)
		assert.Equal(t, tc.want, got)
	}
}
