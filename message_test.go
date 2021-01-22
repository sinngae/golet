package goerrcode

import (
	"fmt"
	"testing"
)

func TestWithMessage(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case 0",
			args:    args{
				err: nil,
				msg: "",
			},
			wantErr: false,
		},
		{
			name:    "case 0",
			args:    args{
				err: fmt.Errorf("err test"),
				msg: "msg test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WithMessage(tt.args.err, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("WithMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMessage(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 0",
			args: args{
				err: WithMessage(fmt.Errorf("err test"), "msg test"),
			},
			want: "msg: msg test, err: err test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Message(tt.args.err); got != tt.want {
				t.Errorf("Message() = %v, want %v", got, tt.want)
			}
		})
	}
}