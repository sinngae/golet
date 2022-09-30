package urctf8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleInit() {

}

func ExampleBits() {
	a := uint64(3)
	b := a << 1
	println(b)
	c := b & a
	println(c)
	// output:

}

func TestRetCode_Uint64(t *testing.T) {
	tests := []struct {
		name      string
		data      []byte
		prefixLen int
		suffixLen int
		want      uint64
	}{
		{
			name:      "case 0",
			data:      []byte{1, 3, 5},
			prefixLen: 0,
			suffixLen: 0,
			want:      0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &RetCode{
				data:      tt.data,
				prefixLen: tt.prefixLen,
				suffixLen: tt.suffixLen,
			}
			got := u.Uint64()
			assert.Equal(t, tt.want, got)
		})
	}
}
