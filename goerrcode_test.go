package goerrcode

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestRetCode_Error(t *testing.T) {
	type fields struct {
		cause   error
		retCode int32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "test case 0: cause=nil",
			fields: fields{
				cause:   nil,
				retCode: 123,
			},
			want: "retcode:123",
		},
		{
			name: "test case 1: cause!=nil",
			fields: fields{
				cause:   errors.New("test"),
				retCode: 123,
			},
			want: "retcode:123, err:test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &withRetCode{
				cause:   tt.fields.cause,
				retCode: tt.fields.retCode,
			}
			if got := rc.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRetCode_RetCode(t *testing.T) {
	type fields struct {
		cause   error
		retCode int32
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		// TODO: Add test cases.
		{
			name: "test case 0",
			fields: fields{
				cause:   nil,
				retCode: 123,
			},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &withRetCode{
				cause:   tt.fields.cause,
				retCode: tt.fields.retCode,
			}
			if got := rc.RetCode(); got != tt.want {
				t.Errorf("RetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithRetCode(t *testing.T) {
	type args struct {
		err     error
		retCode int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test case 0",
			args: args{
				err:     errors.New("test"),
				retCode: 123,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WithRetCode(tt.args.err, tt.args.retCode); (err != nil) != tt.wantErr {
				t.Errorf("WithRetCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// usage
func TestRetCode(t *testing.T) {
	err := WithRetCode(errors.New("test"), 123)
	err = WithRetCode(err, 1234)
	err = errors.WithMessage(err, "this is a err")
	fmt.Println(err)
	fmt.Println(RetCode(err))
}
