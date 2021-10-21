package string

import "strings"

func RuneStrip(src string, limit int) string {
	idx := 0
	return strings.Map(func(r rune) rune {
		idx++
		if idx > limit {
			return -1
		}
		return r
	}, src)
}

func SubRunes(src string, start, end int) string {
	idx := 0
	return strings.Map(func(r rune) rune {
		if idx >= start && idx < end {
			idx++
			return r
		}
		idx++
		return -1
	}, src)
}
