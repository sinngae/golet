package utility

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

type cpSrc struct {
	VInt int32
	VStr string

	VIntBig int32
	VIntSml int64
	VIntFlt int

	VIntUint int
	VUintInt uint

	VFltInt float32

	VFltBig float32
	VFltSml float64

	VStrPtr string
	VPtrStr *string

	VIntStr int
	VStrInt string

	VSrcSpecial int
}

type cpDest struct {
	VInt int32
	VStr string

	VIntBig int64
	VIntSml int32
	VIntFlt float32

	VIntUint int
	VUintInt uint

	VFltInt int

	VFltBig float64
	VFltSml float32

	VStrPtr *string
	VPtrStr string

	VIntStr string
	VStrInt int

	VDestSpecial int
}

func TestMustCopy(t *testing.T) {
	type args struct {
		src  *cpSrc
		want *cpDest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case 0",
			args: args{
				src: &cpSrc{
					VInt:        123,
					VStr:        "test",
					VIntBig:     123,
					VIntSml:     999999999999999,
					VIntFlt:     1000000000001,
					VIntUint:    -123,
					VUintInt:    123,
					VFltInt:     1000000000.123,
					VFltBig:     1.2121212,
					VFltSml:     1.000000000000001,
					VStrPtr:     "test",
					VPtrStr:     proto.String("test"),
					VIntStr:     0123,
					VStrInt:     "234234",
					VSrcSpecial: 234,
				},
				want: &cpDest{},
			},
		},
	}
	for _, tt := range tests {
		out := &cpDest{}
		t.Run(tt.name, func(t *testing.T) {
			MustCopy(out, tt.args.src)
			fmt.Println(out)
		})
	}
}
