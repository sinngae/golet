package bignumber

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func chenji(num1, num2 string) string {
	num1List, num2List := []byte(num1), []byte(num2)
	data := make([]byte, len(num1List)+len(num2List)+1)

	idx1, idx2 := len(num1List)-1, len(num2List)-1
	goNext := []byte{'0'}
	for base1 := 0; base1 < len(num1List); base1++ {
		for base2 := 0; base2 < len(num2List); base2++ {
			base := base1 + base2
			idx := len(data) - base
			res := chengji3(num1List[idx1], num2List[idx2])
			res = add(res, goNext)
			for {
				res = add(res, []byte{data[idx]})
				data[idx] = res[len(res)-1]
				if len(res) == 1 {
					break
				}
				res = res[0 : len(res)-1]
				idx++
			}
			idx1--
			idx2--
		}
	}

	for {
		if data[0] != 0 {
			break
		}
		data = data[1:]
	}

	return string(data)
}

func chengji2(num []byte, val byte) []byte {
	data := make([]byte, 0)
	goNext := []byte{'0'}
	for idx := len(num) - 1; idx >= 0; idx-- {
		res := chengji3(num[idx], val)
		res = add(res, goNext)
		data = append(data, res[len(res)-1])
		if len(res) > 1 {
			goNext = res[0:1]
		} else {
			goNext = []byte{'0'}
		}
	}
	if goNext[0] != '0' {
		data = append(data, goNext[0])
	}
	sort.Slice(data, func(i, j int) bool {
		return i < j
	})
	return data
}

func add(val1, val2 []byte) []byte {
	num1, _ := strconv.Atoi(string(val1))
	num2, _ := strconv.Atoi(string(val2))
	res := num1 + num2 // >= 0; <= 18
	bytes := []byte(strconv.Itoa(res))
	return bytes
}

func chengji3(val1, val2 byte) []byte {
	res := toInt(val1) * toInt(val2)
	bytes := []byte(strconv.Itoa(res)) // >= 0; < 100
	return bytes
}

func toInt(val byte) int {
	return int(val - '0')
}

func TestChenji(t *testing.T) {
	testcases := []struct {
		name string
		num1 string
		num2 string
		want string
	}{
		{"case 1", "123", "456", "56088"},
		{"case 1", "999999999", "99999999", "56088"},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := chenji(tc.num1, tc.num2)
			assert.Equal(t, tc.want, got)
		})
	}
}
