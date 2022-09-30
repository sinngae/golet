package go_std

import (
	"fmt"
	"sort"
	"testing"
)

func Test199(t *testing.T) {
	data := make(map[int]int)
	total := 199
	for idx := 0; idx < total; idx++ {
		val := idx*100000 + 199
		res := val % 199
		if v, ok := data[res]; ok {
			data[res] = v + 1
		} else {
			data[res] = 1
		}
	}

	type result struct {
		val     int
		cnt     int
		percent float64
	}
	list := make([]result, 0, len(data))
	for k, v := range data {
		percent := float64(v) / float64(total)
		list = append(list, result{
			val:     k,
			cnt:     v,
			percent: percent,
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].val < list[j].val
	})
	res := fmt.Sprintf("%+v", list)
	println(res)
}

func Test199big(t *testing.T) {
	data := make(map[int]int)
	total := 1000000
	for idx := 0; idx < total; idx++ {
		val := idx*1000 + 216
		res := val % 199
		if v, ok := data[res]; ok {
			data[res] = v + 1
		} else {
			data[res] = 1
		}
	}

	type result struct {
		val     int
		cnt     int
		percent float64
	}
	list := make([]result, 0, len(data))
	for k, v := range data {
		percent := float64(v) / float64(total)
		list = append(list, result{
			val:     k,
			cnt:     v,
			percent: percent,
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].val < list[j].val
	})
	res := fmt.Sprintf("%+v", list)
	println(res)
}
