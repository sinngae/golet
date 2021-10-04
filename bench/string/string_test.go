package string

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestTrimFunc(t *testing.T) {
	testcases := []struct {
		val  string
		want string
	}{
		{"1a2b3c4", ""},
	}
	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
			got := strings.TrimFunc(tc.val, func(r rune) bool { return unicode.IsDigit(r) })
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMap(t *testing.T) {
	testcases := []struct {
		val  string
		want string
	}{
		{"1a2b3c4", ""},
	}
	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
			got := strings.Map(func(r rune) rune {
				if unicode.IsDigit(r) {
					return -1
				}
				return r
			}, tc.val)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestTrim(t *testing.T) {
	testcases := []struct {
		val   string
		strip string
		want  string
	}{
		{"abcefabcefabc", "abc", ""},
	}
	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
			got := strings.Trim(tc.val, tc.strip)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSprintf(t *testing.T) {
	testcases := []struct {
		fmt  string
		val  float64
		want string
	}{
		{"%.0f", 123.123, "123"},
		{"%.0f", 123.512, "123"},
		{"%.0f", 123.5, "123"},
		{"%.0f", 123.0, "123"},
	}
	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
			got := math.Floor(tc.val)
			assert.Equal(t, tc.want, got)
		})
	}
}
