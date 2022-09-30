package rc_int

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRetCode(t *testing.T) {
	tests := []struct {
		name string
		pre  uint64
		suf  uint64
		want uint64
	}{
		{"abc", 0, 0, 0},
		{"abc1", 0, 1, 1},
		{"abc1", 1, 1, 4294967297},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRetCode(tt.pre, tt.suf)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		val   uint64
		want  uint64
		want1 uint64
	}{
		{"case0", 12421342394, 2, 3831407802},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Parse(tt.val)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
