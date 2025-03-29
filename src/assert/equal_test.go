package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		name string

		expected interface{}
		actual   interface{}

		want  string
		want1 bool
	}{
		{"case 1", "abc", "def", "", false},
		{"case 1", "abc", 123, "", false},
		{"case 1", struct{ val int }{123}, struct{ val string }{"abc"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := assert.Equal(t, tt.expected, tt.actual)
			assert.NotEqual(t, tt.want, got)
		})
	}
}
