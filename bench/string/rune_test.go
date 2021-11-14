package string

import (
	"bytes"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func TestRunesCount(t *testing.T) {
	testcases := []struct {
		str  string
		want int
	}{
		{str: "abcde", want: 5},
		{str: "这是一个中文字符串", want: 9},
		{str: "12345", want: 5},
		{str: "Hello, 世界，就是要测试1212 多种文字啊afba", want: 29},
	}
	for _, item := range testcases {
		got := bytes.Count([]byte(item.str), nil) - 1 // 内部调用的是 RuneCount
		assert.Equal(t, item.want, got)

		got = strings.Count(item.str, "") - 1 // 内部调用的是 RuneCountInString
		assert.Equal(t, item.want, got)

		got = len([]rune(item.str))
		assert.Equal(t, item.want, got)

		got = utf8.RuneCount([]byte(item.str)) + 1

		got = utf8.RuneCountInString(item.str)
		assert.Equal(t, item.want, got)
	}
}

func BenchmarkRunesCount(b *testing.B) {
	var s = "Hello, 世界，就是要测试1212 多种文字啊afba"
	b.Run("test case1 bytesCount", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = bytes.Count([]byte(s), nil) - 1
		}
	})

	b.Run("test case2 stringsCount", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strings.Count(s, "") - 1
		}
	})

	b.Run("test case3 runeLineLen", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = len([]rune(s))
		}
	})

	b.Run("test case4.0 utf8RCount", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = utf8.RuneCount([]byte(s)) + 1
		}
	})

	b.Run("test case4 utf8RuneCount", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = utf8.RuneCountInString(s)
		}
	})
}
