package go_std

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelBit(t *testing.T) {
	testcases := []struct {
		name string
		flag uint32
		mask uint32
		want uint64
	}{
		{"hilife channel flag test", 67436560, 1 << 16, 1 << 63},    // 65536
		{"hilife channel flag staging", 67371026, 1 << 16, 1 << 16}, // 0
		{"hilife channel flag live", 67436544, 1 << 16, 1 << 16},    // 65536
	}
	for _, item := range testcases {
		t.Run("", func(t *testing.T) {
			got := item.flag & item.mask
			assert.Equal(t, item.want, got)
		})
	}
}
