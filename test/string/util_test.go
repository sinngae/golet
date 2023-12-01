package string

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestStrCut(t *testing.T) {
	tests := []struct {
		src   string
		limit int
		want  string
	}{
		{"abcdefg", 3, "abc"},
		{"abcdefg", 5, "abcde"},
		{"abc", 5, "abc"},
		{"这是一个中文字符串", 3, "这是一"},
		{"这是一个中文字符串", 5, "这是一个中"},
		{"这是一", 5, "这是一"},
	}
	for _, tt := range tests {
		t.Run(tt.src, func(t *testing.T) {
			got := RuneStrip(tt.src, tt.limit)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSubRunes(t *testing.T) {
	tests := []struct {
		src   string
		start int
		end   int
		want  string
	}{
		{"abcdefg", 0, 3, "abc"},
		{"abcdefg", 2, 5, "cde"},
		{"abc", 0, 5, "abc"},
		{"这是一个中文字符串", 0, 3, "这是一"},
		{"这是一个中文字符串", 2, 5, "一个中"},
		{"这是一", 0, 5, "这是一"},
	}
	for _, tt := range tests {
		t.Run(tt.src, func(t *testing.T) {
			got := SubRunes(tt.src, tt.start, tt.end)
			assert.Equal(t, tt.want, got)
		})
	}
}
