package algo

import (
	"fmt"
	"testing"
)

func threeSum2(nums []int) [][]int {
	sNums := sortInts(nums)

	var numPre int

	dest := make([][]int, 0)
	maxIdx := len(sNums) - 2
	for i := 0; i < maxIdx; i++ {
		if i != 0 && nums[i] == numPre {
			continue
		}
		numPre = sNums[i]

		rIdx, lIdx := len(sNums)-1, i+1
		// rPre, lPre := sNums[rIdx], sNums[lIdx]
		//ok := false
		for rIdx > lIdx {
			val := sNums[i] + sNums[rIdx] + sNums[lIdx]
			switch {
			case val == 0:
				dest = append(dest, []int{sNums[i], sNums[lIdx], sNums[rIdx]})
				for rIdx > lIdx {
					if sNums[rIdx] != sNums[rIdx-1] {
						break
					}
					rIdx--
				}
				rIdx--
				for rIdx > lIdx {
					if sNums[lIdx] != sNums[lIdx+1] {
						break
					}
					lIdx++
				}
				lIdx++
			case val > 0:
				for rIdx > lIdx {
					if sNums[rIdx] != sNums[rIdx-1] {
						break
					}
					rIdx--
				}
				rIdx--
			case val < 0:
				for rIdx > lIdx {
					if sNums[lIdx] != sNums[lIdx+1] {
						break
					}
					lIdx++
				}
				lIdx++
			}
		}
	}
	return dest
}

func TestThreeSum2(t *testing.T) {
	data := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum2(data)
	fmt.Println(res)
}

func twoSum(nums []int, sum int) [][]int {
	var numPre int
	dest := make([][]int, 0)
	max := len(nums) - 1
	for i := 0; i < max; i++ {
		if i != 0 && nums[i] == numPre {
			continue
		}
		numPre = nums[i]
		if nums[i] > sum {
			break
		}

		val := sum - nums[i]
		if val > nums[max] {
			continue
		}
		if findVal(nums[i+1:], val) {
			dest = append(dest, []int{nums[i], val})
		}
	}
	return dest
}

func findVal(nums []int, val int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		if nums[0] == val {
			return true
		}
		return false
	}

	idxMid := len(nums) / 2
	num := nums[idxMid]
	switch {
	case num == val:
		return true
	case num > val:
		return findVal(nums[0:idxMid], val)
	}
	return findVal(nums[idxMid+1:], val)
}
